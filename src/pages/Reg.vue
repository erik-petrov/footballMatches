<script>
import { reactive, computed } from "vue";
import useValidate from "@vuelidate/core";
import { required, email, minLength, sameAs, not, minValue, helpers } from "@vuelidate/validators";
import axios from "axios";
import router from "../router";
export default {
  name: "Reg",
  data() {
    return{
      selected: '',
      usedEmails: []
    }
  },
  setup(){
    const state = reactive({
      _email: "",
      _name: "",
      _sname:"",
      _pswd: "",
      _pswdR: "",
      _date: "",
    })
    const rules = computed(() =>{
      return{
        _name: {required, not: not(sameAs("Олег"))},
        _pswd: {required, minLength: minLength(12)},
        _date: {required, minValue: value => value < new Date().toISOString()},
        _email: {required, email},
        _pswdR: {required, sameAs: sameAs(state._pswd)},
      }
    })
    const v$ = useValidate(rules, state)
    return {state, v$}
  },
  methods: {
    submitForm($e) {
      this.v$.$touch();
      if (this.v$.$error) {
        return
      }
      let genderText = this.$refs["gender"]
      if(genderText.style.visibility == "visible"){
        this.selected = genderText.value
      }
      try {
        let data = {name: this.state._name, sname: this.state._sname, email: this.state._email, password: this.state._pswd, date: this.state._date, gender: this.selected}
        axios.post("http://localhost:9000/reg", data).catch(err => console.log(err))
        router.push({name: 'Login'})
      } catch (e){
        console.log("error")
        console.log(e.message)
      }
    },
    detectChange(e){
      if(e.target.value == "c"){
        this.$refs["gender"].style.visibility = "visible";
      }else{
        this.$refs["gender"].style.visibility = "hidden";
      }
    },
  }
}

</script>

<template>
  <div class="signup-form">
    <form @submit.prevent="submitForm" method="post">

      <div>
          <span v-if="v$._name.$error">Required. Олегам тоже нельзя.</span><br>
        <label for="name">First name</label><br>
        <input class="form-control" id="name" name="name" type="text" v-model="state._name" required/>
      </div>

      <div>
        <label for="sname">Last name</label><br>
        <input class="form-control" id="sname" name="sname" type="text" v-model="state._sname" required/>
      </div>
      <br>
      <div>
        <span v-if="v$._email.$error">Wrong email format or it is taken</span><br>
        <label for="email">Email</label><br>
        <input class="form-control" id="email" name="email" type="email" v-model="state._email" required/>
      </div>
      <br>
      <div>
        <span v-if="v$._pswd.$error">Passwords minimum length is 12 characters</span><br>
        <label for="pswd">Password</label><br>
        <input class="form-control" id="pswd" name="pswd" type="password" v-model="state._pswd" required/>
      </div>
      <div>
        <span v-if="v$._pswdR.$error">Passwords don't match</span><br>
        <label for="pswdR">Repeat password</label><br>
        <input class="form-control" id="pswdR" name="pswdR" type="password" v-model="state._pswdR" required/>
      </div><br>

      <label for="gender">Gender</label><br>
      <select name="gender" class=".form-control" id="gender" v-on:change="detectChange($event)" v-model="selected" required>
        <option disabled value="">Choose your gender..</option>
        <option value="Male">Male</option>
        <option value="Female">Female</option>
        <option value="c">Other..</option>
      </select><br>
      <input class=".form-control" name="gender" style="visibility: hidden;" ref="gender" placeholder="Впишите свой гендер" type="text" /><br>

      <span v-if="v$._date.$error">You can't be born in the future</span><br>
      <label for="date">Birthday</label><br>
      <input class="form-control" id="date" name="date" type="date" v-model="state._date" required/><br><br>


      <button class="btn--form" type="submit" >Register</button>

    </form>
  </div>
</template>



<style scoped>
[type="submit"] {transition: all .25s ease-in;}
.btn--form {
  padding: .5rem 2.5rem;
  font-size: .95rem;
  font-weight: 600;
  text-transform: uppercase;
  color: #fff;
  background: #111;
}
.form-control {
  background-color: transparent;
  border-top: 0;
  border-right: 0;
  border-left: 0;
  border-radius: 0;
  color: white;
}
.form-control:active{border-color: #ffffff;}
label {
  font-size: .85rem;
  text-transform: uppercase;
  color: #777777;
}
.signup-form{
  max-width: 50%;
  margin: auto;
}
form{
  align-content: center;
  display: block;
}
</style>
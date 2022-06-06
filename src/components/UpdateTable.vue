<template>
  <button @click="openModal">Change user data</button>
  <vue-final-modal v-model="showModal" name="modal" classes="modal-container" content-class="modal-content">
    <div id="update-form">
      <span class="modal__title">
        <slot>Enter new data</slot>
      </span>
      <form @submit.prevent="submitForm" method="post">
        <div>
          <span v-if="v$._name.$error">Олегам нельзя.</span><br>
          <label for="name">Имя</label><br>
          <input class="form-control" id="name" name="name" type="text" v-model="state._name" required/>
        </div>

        <div>
          <label for="sname">Lastname</label><br>
          <input class="form-control" id="sname" name="sname" type="text" v-model="state._sname" required/>
        </div>
        <br>

        <div>
          <span v-if="v$._email.$error">Wrong email format</span><br>
          <label for="email">Email</label><br>
          <input class="form-control" id="email" name="email" type="email" v-model="state._email" required/>
        </div>
        <br>

        <div>
          <span v-if="v$._pswdOG.$error">Wrong password</span><br>
          <label for="pswdOG">Current password</label><br>
          <input class="form-control" id="pswdOG" name="pswdOG" type="password" v-model="state._pswdOG" required/>
        </div>

        <div class="change-password">
          <div>
            <span v-if="v$._pswd.$error">Passwords minimum length is 12 characters</span><br>
            <label for="pswd">New password</label><br>
            <input class="form-control" id="pswd" name="pswd" type="password" v-model="state._pswd"/>
          </div>
          <div>
            <span v-if="v$._pswdR.$error">Passwords don't match</span><br>
            <label for="pswdR">Repeat password</label><br>
            <input class="form-control" id="pswdR" name="pswdR" type="password" v-model="state._pswdR"/>
          </div><br>
        </div>

        <label for="gender">Gender</label><br>
        <input class=".form-control" name="gender" ref="gender" placeholder="Впишите свой гендер" v-model="state._gender" type="text" /><br>

        <span v-if="v$._date.$error">You can't be born in the future</span><br>
        <label for="date">Birthday</label><br>
        <input class="form-control" id="date" name="date" type="date" v-model="state._date" required/><br><br>


        <button class="btn--form" type="submit" >Save</button>

      </form>
      <button class="modal__close" @click="closeModal">X</button>
    </div>
  </vue-final-modal>
</template>

<script>
import axios from "axios";
import {email, helpers, minLength, not, required, sameAs} from "@vuelidate/validators";
import {computed, reactive, ref} from "vue";
import useValidate from "@vuelidate/core";
import router from "../router";
const { withAsync } = helpers
export default{
  name: 'UpdateTable',
  data: () => ({
    showModal: false,
    name: sessionStorage.getItem("name").split(':')[0],
    lname: sessionStorage.getItem("name").split(':')[1],
    created_at: sessionStorage.getItem("created_at"),
    email: sessionStorage.getItem("email"),
    bday: sessionStorage.getItem("bday").split('T')[0],
    gender: sessionStorage.getItem("gender")
  }),
  setup() {
    const state = reactive({
      _email: "",
      _name: "",
      _pswdOG: "",
      _sname:"",
      _pswd: "",
      _pswdR: "",
      _gender: "",
      _date: "",
    })
    const rules = computed(() =>{
      return{
        _name: {not: not(sameAs("Олег"))},
        _pswd: {minLength: minLength(12)},
        _date: {minValue: value => value < new Date().toISOString()},
        _email: {email},
        _pswdR: {sameAs: sameAs(state._pswd)},
        _pswdOG: {required}
      }
    })
    const v$ = useValidate(rules, state);
    return {
      state,
      v$
    }
  },
  methods: {
    openModal() {
      this.state._name = this.name.replace(/\s+/g, '');
      this.state._sname = this.lname;
      this.state._email = this.email;
      this.state._gender = this.gender;
      this.state._date = this.bday;
      this.$vfm.show('modal')
    },
    closeModal() {
      this.$vfm.toggle('modal')
    },
    submitForm() {
      this.v$.$touch();
      if (this.v$.$error) {
        return
      }
      let data = {email: this.email, password: this.state._pswdOG}
      axios.post("http://localhost:9000/login", data)
          .then((res) => {
          })
          .catch(err => {console.log(err); return;})

      let id = sessionStorage.getItem("id")
      //TODO: fill session storage after restart
      //TODO: make table and filter maybe?
      if(this.state._pswd != ""){
        try{
          let data = {id: id, pswd: this.state._pswd}
          axios.post("http://localhost:9000/pswdChange", data).catch(err => console.log(err))
        }
        catch (e){
          console.log(e.message)
          return
        }
      }
      try {
        let myDate = new Date(this.state._date).toISOString().split('T')[0]
        let data = {id: id, name: this.state._name, sname: this.state._sname, email: this.state._email, password: this.state._pswdOG, date: myDate, gender: this.state._gender}
        axios.post("http://localhost:9000/edit", data)
            .then((res) => {
              const userData = res.data
              sessionStorage.setItem("loggedIn", "true")
              sessionStorage.setItem("id", userData.id)
              sessionStorage.setItem("name", userData.name)
              sessionStorage.setItem("email", userData.email)
              sessionStorage.setItem("created_at", userData.created_at)
              sessionStorage.setItem("bday", userData.birthday)
              sessionStorage.setItem("gender", userData.gender)
              router.go()
            })
            .catch(err => console.log(err))
      } catch (e){
        console.log(e.message)
      }
    },
  }
}
</script>

<style scoped>
::v-deep(.modal-container) {
  position: absolute !important;
  display: flex;
  justify-content: center;
  align-items: center;
}
::v-deep(.modal-content) {
  position: relative;
  display: flex;
  flex-direction: column;
  max-height: 90%;
  margin: 0 1rem;
  padding: 1rem;
  border: 1px solid #e2e8f0;
  border-radius: 0.25rem;
  background: #1d1d1d;
}

.modal__title {
  margin: 0 2rem 0 0;
  font-size: 1.5rem;
  font-weight: 700;
}

.modal__close {
  position: absolute;
  top: -0.5rem;
  right: -0.5rem;
}
vue-final-modal{
  width: 90%;
}
</style>
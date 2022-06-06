<template>
  <div class="signup-form">
    <span v-if="error" style="background-color: #e76868; border: 1px solid red; color: #620202;">Wrong credentials.</span><br>
    <form @submit.prevent="submitForm" method="post">
      <div>
        <span v-if="v$._email.$error">Wrong email </span><br>
        <label for="email">Email</label><br>
        <input class="form-control" id="email" name="email" type="email" v-model="state._email" required/>
      </div>
      <br>
      <div>
        <label for="pswd">Password</label><br>
        <input class="form-control" id="pswd" name="pswd" type="password" v-model="state._pswd" required/>
      </div>
      <br>

      <button class="btn--form" type="submit" >Войти</button>

    </form>
  </div>
</template>

<script>
import { reactive, computed } from "vue";
import useValidate from "@vuelidate/core";
import { required, email} from "@vuelidate/validators";
import axios from "axios";
import router from "../router";

export default {
  name: "Login",
  setup() {
    const state = reactive({
      _email: "",
      _pswd: "",
    })
    const rules = computed(() => {
      return {
        _pswd: {required},
        _email: {required, email},
      }
    })
    const v$ = useValidate(rules, state)
    return {state, v$}
  },
  data() {
    return{
      error: false,
    }
  },
  methods: {
    submitForm() {
      this.v$.$touch();
      if (this.v$.$error) {
        return
      }
      try {
        let data = {email: this.state._email, password: this.state._pswd}
        axios.post("http://localhost:9000/login", data)
            .then((res) => {
              const userData = res.data
              sessionStorage.setItem("loggedIn", "true")
              sessionStorage.setItem("id", userData.id)
              sessionStorage.setItem("name", userData.name)
              sessionStorage.setItem("email", userData.email)
              sessionStorage.setItem("bday", userData.birthday)
              sessionStorage.setItem("createdAt", userData.created_at)
              sessionStorage.setItem("gender", userData.gender)
              this.emitter.emit("logged")
              router.push({name: "UserHome"})
            })
            .catch(err => {this.error = true; throw err})
      } catch (e){
        console.log(e.message)
      }
    },
  }
}
</script>

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
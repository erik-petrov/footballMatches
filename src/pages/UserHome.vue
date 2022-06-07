<script setup>
import day from 'dayjs'
import UpdateTable from "../components/UpdateTable.vue";
import MatchesTable from "../components/MatchesTable.vue";
import relativeTime from 'dayjs/plugin/relativeTime.js';
import utc from 'dayjs/plugin/utc.js';
import tz from 'dayjs/plugin/timezone'

day.extend(tz)
day.extend(relativeTime)
day.extend(utc)
</script>
<script>
export default {
  name: "UserHome",
  data(){
    return{
      userName: "",
      created_at: "",
      email: "",
      bday: "",
      gender: "",
      usedEmails: [],
      selected: "",
      showModal: false,
    }
  },

  mounted() {
    let nameArr = sessionStorage.getItem("name").split(':')
    this.email = sessionStorage.getItem("email")
    this.userName = nameArr[0] + " " + nameArr[1]
    this.created_at = sessionStorage.getItem("createdAt")
    this.bday = sessionStorage.getItem("bday")
    this.gender = sessionStorage.getItem("gender")
  },
  methods:{
    getAge(dateString){
      let today = new Date();
      let birthDate = new Date(dateString.split('T')[0]);
      console.log(dateString.split('T')[0])
      let age = today.getFullYear() - birthDate.getFullYear();
      let m = today.getMonth() - birthDate.getMonth();
      if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
        age--;
      }
      return age;
    },
    detectChange(e){
      if(e.target.value == "c"){
        this.$refs["gender"].style.visibility = "visible";
      }else{
        this.$refs["gender"].style.visibility = "hidden";
      }
    }
  }
}
</script>
<template>
  <div class="split">
    <div class="left">
      <div class="centered">
        <h2>Welcome {{userName}}!</h2>
        <div class="date-box">
          <p>You've been with us for: {{day(this.created_at).subtract('3', 'h').fromNow(true)}}</p>
          <p>You are: {{getAge(this.bday)}} years old</p>
          <p>Your gender: {{this.gender}}</p>
        </div>
      </div>
      <div class="updateTable">
        <UpdateTable />
      </div>
    </div>
    <div class="right">
      <matches-table></matches-table>
    </div>
  </div>
</template>



<style scoped>
.split {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
}
.left {
  flex: 1;
}

.right {
  flex: 3;
  margin: 10% auto auto;
  -moz-transform: scale(1.3);
  zoom: 1.3%;
  justify-content: space-around;
}

.centered {
  top: 50%;
  left: 50%;
  transform: translate(-50%, 0%);
  text-align: center;
}

.updateTable{
  text-align: center;
}
</style>
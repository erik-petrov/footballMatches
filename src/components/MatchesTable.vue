

<template>
  <div class="table">
    <table>
      <thead class="tableHead">
      <tr>
        <th>Date</th>
        <th>Team 1</th>
        <th>Score</th>
        <th>Team 2</th>
        <th>Length</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="match in matches" :key="match.id" :class="getClass(`${match.date}`)">
        <td class="date">{{ match.date.split('T')[0]}}</td>
        <td class="team">{{ match.team1 }}</td>
        <td class="scores">{{ match.score }}</td>
        <td class="team">{{ match.team2 }}</td>
        <td class="length">{{ match.length }}</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import axios from 'axios'

export default{
  props: ['id'],
  data() {
    return{
      matches: []
    }
  },
  mounted() {
    /*fetch('http://localhost:9000/matches')
        .then(res => res.json())
        .then(data => this.matches = data)
        .catch(err => console.log(err.message))*/
    axios('http://localhost:9000/matches')
        .then(data => {
          this.matches = data.data
        })
        .catch(err => console.log(err.message))
  },
  methods: {
    getClass(a){
      let match = new Date(a)
      let timeDif = match.getTime() - new Date();
      let dayDif = timeDif / (1000 * 3600 * 24);
      let amount = Math.floor(dayDif)
      console.log(amount)
      if (amount == -1) {
        this.class="now"
        return this.class
      }
      if (amount > 0) {
        this.class = "will"
        return this.class
      }
      else {
        this.class = "was"
        return this.class
      }
    }
  }
}
</script>
<style scoped>
.now{
  background-color: #366b66;
}

.will{
  background-color: #3c4b30;
}

.table{
  width: 100%;
}
.scores:not(:empty){
  padding: 2px;
  text-align: center;
  border: 1px #2ae5d5 solid;
}
.date{
  padding-right: 10px;
}
.tableHead{
  border: 1px red solid;
}
.length{
  padding: 2px;
}

.table th{
  border-bottom: 1px hsla(160, 100%, 37%, 1) solid;
}

.date, .team, .length{
  padding-left: 50px;
  padding-right: 50px;
}
</style>
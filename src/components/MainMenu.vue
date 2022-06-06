<template>
  <div id="menu">
    <div id="nav-bar">
      <router-link id="logoA" :to="{name: 'Home'}">
        <svg id="logo" xmlns="http://www.w3.org/2000/svg" viewBox="-2500 -2500 5000 5000"><g stroke="#000" stroke-width="24"><circle fill="#fff" r="2376"/><path fill="none" d="m-1643-1716 155 158m-550 2364c231 231 538 195 826 202m-524-2040C-2227-681-2346 32-2328 28m1216-1008c-51 373 84 783 364 1220m-107-2289c157-157 466-267 873-329m-528 4112c-50 132-37 315-8 510m62-3883c282 32 792 74 1196 303M336 1308c310 173 649 247 1060 180M1056-520C814-186 522 125 184 416m1109-2119c-111-207-296-375-499-534M1940-956c100 3 197 44 290 141m-438 495c158 297 181 718 204 1140"/></g><path d="M-1624-1700c243-153 498-303 856-424 141 117 253 307 372 492-288 275-562 544-724 756-274-25-410-2-740-60 3-244 84-499 236-764zm2904-40c271 248 537 498 724 788-55 262-105 553-180 704-234-35-536-125-820-200-138-357-231-625-340-924 210-156 417-296 616-368zm-3273 3033A2376 2376 0 0 1-2371-99l59-7c54 342 124 674 311 928-36 179-2 323 51 458zM-796 168c365 60 717 120 1060 180 106 333 120 667 156 1000-263 218-625 287-944 420-372-240-523-508-736-768 122-281 257-561 464-832zm3013 678a2376 2376 0 0 1-925 1147l-116-5c84-127 114-297 118-488 232-111 464-463 696-772 86 30 159 72 227 118zM-70 2373a2376 2376 0 0 1-993-251c199 74 367 143 542 83 53 75 176 134 451 168z"/></svg>
      </router-link>
      <ul id="mainMenu">
        <div v-if="!this.isLogged" style="display: inline">
          <li><router-link :to="{name: 'Login'}">Log in</router-link></li>
          <li><router-link :to="{name: 'Reg'}">Sign up</router-link></li>
        </div>
        <div v-else style="display: inline">
          <li><router-link :to="{path: 'UserHome'}">Matches</router-link></li>
          <li><a class="pointerThing" @click.prevent="signOut">Log out</a></li>
        </div>
        <li><div class="dropdown"><fa icon="list" id="list"/><div class="dropdown-content">
          <a href="#">I</a>
          <a href="#">Hate</a>
          <a href="#">Css</a>
        </div></div></li>
      </ul>
    </div>
  </div>
</template>

<script>
import router from "../router";

export default {
  name: "MainMenu",
  data(){
    return{
      isLogged: this.checkLoggedState()
    }
  },
  created () {
     this.emitter.on('logged', () => {
      this.isLogged = this.checkLoggedState()
    })
  },
  methods: {
    checkLoggedState(){
      if(sessionStorage.getItem("loggedIn") === "true"){
        return true
      }
      else{
        return false
      }
    },
    signOut(){
      sessionStorage.clear()
      sessionStorage.setItem("loggedIn", "false")
      this.emitter.emit("logged")
      router.push({name: 'Login'})
    }
  }
}
</script>

<style scoped>
.dropdown{
  position: relative;
  padding-top: 25px;
  padding-left: 10px;
  padding-right: 10px;
  float: right;
}

.pointerThing:hover{
  cursor: pointer;
}

.dropdown-content {

  right: 0px;
  display: none;
  position: absolute;
  min-width: 160px;
  padding: 6px 8px !important;
  z-index: 1;
}

.dropdown-content a {
  transition: 0.4s;
  text-decoration: none;
  display: block;
}

.dropdown:hover .dropdown-content {
  display: block;
}

#menu{
  padding-bottom: 10px;
}

#nav-bar {
  background-color: #2c2c2c;
  height: 10%;
  width: 100%;
  text-align:right;
  box-shadow: 11px 2px 8px 2px #282828;
}

#list{
  color: hsla(160, 100%, 37%, 1);
  height: 25px;
}
#logoA:hover{
  background-color: #2c2c2c;
}
#logoA{
  display: inline;
  text-align: left !important;
}

#logo{
  height: 70px;
  float:left;
}
#nav-bar ul {
  padding: 0px;
  margin: 0px;
  text-align: center;
  display:inline-block;
  vertical-align:top;
}

#nav-bar li {
  list-style-type: none;
  padding: 0px;
  height: 24px;
  margin-top: 4px;
  margin-bottom: 4px;
  display: inline;
}
#nav-bar a{
  /*color: black;*/
  font-size: 16px;
  font-family: "Trebuchet MS", Arial, Helvetica, sans-serif;
  text-decoration: none;
  line-height: 70px;
  padding: 5px 15px;
  opacity: 0.7;
}

</style>
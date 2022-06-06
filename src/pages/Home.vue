<template>
  <div class="content">
    <audio id="saul" loop autoplay>
      <source src="../assets/saul.mp3" type="audio/mpeg">
    </audio>
    <h1>Добро пожаловать на с̶̤̫͊а̵̻͈̓̚й̶̟͖͑͘т̶̻͎͌ ̸͇̣͆̌с̷͉̕ ̷̭̳̿͋ф̷͎͗̀у̸͓͗̀т̸̲̄̕б̷͖͕̏̋о̵̦̻͛̇л̶̠̠͂̚ь̸͚͉̈н̷͇͆͋ы̶̻̮̎͠м̶͓̤͠и̸̻͖̐̏ ̷̧̟͋м̵̡̣̄а̴̯̂̏т̵̳̪͆ч̷̬̥͝а̶͎̰͒̍м̵̻͒̇ͅи̸̬́</h1>

  </div>
  <div class="imgbox">
  <img class="center-fit" src="https://c.tenor.com/1y8zDc-ll-EAAAAd/3d-saul-saul-goodman.gif" alt="google saul goodman 3d gif"></div>
</template>

<script>
export default {
  name: "Home",
  mounted() {
    window.addEventListener('load', function () {
      var audioCtx = new (window.AudioContext || window.webkitAudioContext)();
      var source = audioCtx.createBufferSource();
      var xhr = new XMLHttpRequest();
      xhr.open('GET', 'audio-autoplay.wav');
      xhr.responseType = 'arraybuffer';
      xhr.addEventListener('load', function (r) {
        audioCtx.decodeAudioData(
            xhr.response,
            function (buffer) {
              source.buffer = buffer;
              source.connect(audioCtx.destination);
              source.loop = false;
            });
        source.start(0);
      });
      xhr.send();
    });
    document.getElementById("saul").volume = 0.02;
  }
}
</script>

<style scoped>
.content{
  max-width: 50%;
  margin: auto;
}
.imgbox {
  display: grid;
  height: 100%;
}
.center-fit {
  max-width: 100%;
  max-height: 200vh;
  margin: auto;
}
h1{
  text-align: center;
  margin-bottom: 100px;
}
</style>
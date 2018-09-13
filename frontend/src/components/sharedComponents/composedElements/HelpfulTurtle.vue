<template>
  <div class='HelpfulTurtle'>
    <div id='shell'/>
    <div id='flipper1'/>   
    <div id='flipper2'/>   
    <div id='flipper3'/>   
    <div id='flipper4'/>   
    <div class='hexContainer'>
        <div class="hex"/>
        <div class='info'>
            i
        </div>
    </div>
    <div class='headContainer'>
        <div id='head'/>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelpfulTurtle',
  data: function(){
    return {
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.HelpfulTurtle{
  height: 100%;
  width: 100%;
  font-size: 3px;
  position: relative;
  cursor: pointer;
}

#shell { 
    position: absolute;
    top: 0;
    left: 37.5%;
    width: 80%;
    height: 90%; 
    background: rgb(73, 160, 120);;
    display: block;
    transform: rotate(180deg);
    border-radius: 50% 50% 50% 50% / 60% 60% 40% 40%;
    z-index: 2;
} /* Cleaner, but slightly less support: use "50%" as value */ 

#flipper1{
    position: absolute;
    top: 25%;
    left: 25%;
    z-index: 1;
    transform: rotate(240deg);
    width: 40%; 
    height: 50%; 
    border-radius: 50%; 
    box-shadow: 3em 1em 0 0 #216869;
}

#flipper2{
    position: absolute;
    top: 25%;
    left: 87.5%;
    z-index: 1;
    transform: rotate(120deg) scaleX(-1);
    width: 40%; 
    height: 50%; 
    border-radius: 70%; 
    box-shadow: 3em 1em 0 0 #216869;
}

#flipper3{
    position: absolute;
    top: 50%;
    left: 75%;
    z-index: 1;
    transform: rotate(-45deg);
    width: 20%; 
    height: 60%; 
    border-radius: 50%; 
    box-shadow: 3em 1em 0 0 #216869;
}

#flipper4{
    position: absolute;
    top: 50%;
    left: 60%;
    z-index: 1;
    transform: rotate(45deg) scaleX(-1);
    width: 20%; 
    height: 60%; 
    border-radius: 50%; 
    box-shadow: 3em 1em 0 0 #216869;
}

.hex {
  position: relative;
  margin: 1em auto;
  width: 10em; height: 17.32em;
  border-radius: 1em/.5em;
  background: orange;
  transition: opacity .5s;
}
.hex:before, .hex:after {
  position: absolute;
  width: inherit; height: inherit;
  border-radius: inherit;
  background: inherit;
  content: '';
}
.hex:before {
    -webkit-transform: rotate(60deg);
  transform: rotate(60deg);
}
.hex:after {
    -webkit-transform: rotate(-60deg);
  transform: rotate(-60deg);
}

.hexContainer{
    position: absolute;
    z-index: 20;
    font-size: 0.1rem;
    left: 61.5%;
    top: 10%;
}

.headContainer{
    position: absolute;
    z-index: 1;
    left: 61.5%;
    top: -22.5%;
}

#head{ 
    position: relative; 
    background:  #216869;
    height: 15px; 
    width: 15px; 
    border-radius: 50%
} 

.info{
    position: absolute;
    z-index: 50;
    top: 15%;
    left: 20%;
    font-family: var(--turtleFont);
    font-weight: 900;
    color: #216869;
    font-size: 2rem;
}

</style>

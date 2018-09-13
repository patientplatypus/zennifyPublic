<template>
  <div class='HelpfulTurtle'>
    <div class='circleBorder'/>
    <div class='innerDiamond'>
        <div class='container'>  
            <div class='chevron' id='chevron1'/>
            <div class='chevron' id='chevron2'/>
            <div class='chevron' id='chevron3'/>
            <div class='chevron' id='chevron4'/>
            <div class='chevron' id='chevron5'/>
            <div class='chevron' id='chevron6'/>
            <div class='chevron' id='chevron7'/>
            <div class='chevron' id='chevron8'/>
            <div class='chevron' id='chevron9'/>
        </div>
    </div>
    <div class='outerCircle'>
        <div class='container'>
            <div class='spark'/>
            <div class='spark' id='sparkRotate1'/>
            <div class='spark' id='sparkRotate2'/>  
            <div class='spark' id='sparkRotate3'/>
            <div class='spark' id='sparkRotate4'/>
            <div class='spark' id='sparkRotate5'/>
            <div class='spark' id='sparkRotate6'/>
            <div class='spark' id='sparkRotate7'/>            
            <div class='spark' id='sparkRotate8'/>
            <div class='spark' id='sparkRotate9'/>
            <div class='spark' id='sparkRotate10'/>
            <div class='spark' id='sparkRotate11'/>
            <div class='spark' id='sparkRotate12'/>
            <div class='spark' id='sparkRotate13'/>
            <div class='spark' id='sparkRotate14'/>
            <div class='spark' id='sparkRotate15'/>            
            <div class='spark' id='sparkRotate16'/>
            <div class='spark' id='sparkRotate17'/>
            <div class='spark' id='sparkRotate18'/>            
            <!-- <div class='spark' id='sparkRotate19'/>
            <div class='spark' id='sparkRotate20'/> -->
        </div>
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
  position: relative;
}


.innerDiamond{
    -webkit-clip-path: polygon(50% 0%, 100% 25%, 100% 75%, 50% 100%, 0% 75%, 0% 25%);
    clip-path: polygon(50% 0%, 100% 25%, 100% 75%, 50% 100%, 0% 75%, 0% 25%);
    height: 25%;
    width: 12.5%;
    position: absolute;
    top: 37.5%;
    left: 43.75%;
    transform: rotate(0deg);
    background: #216869;
    z-index: 3;
}

.outerOctagon{
    -webkit-clip-path: polygon(30% 0%, 70% 0%, 100% 30%, 100% 70%, 70% 100%, 30% 100%, 0% 70%, 0% 30%);
    clip-path: polygon(30% 0%, 70% 0%, 100% 30%, 100% 70%, 70% 100%, 30% 100%, 0% 70%, 0% 30%);
    position: absolute;
    z-index: 1;
    background: rgb(223, 200, 120);
    height: 50%;
    width: 50%;
    top: 25%;
    left: 25%;
}

.octagonBorder{
    -webkit-clip-path: polygon(30% 0%, 70% 0%, 100% 30%, 100% 70%, 70% 100%, 30% 100%, 0% 70%, 0% 30%);
    clip-path: polygon(30% 0%, 70% 0%, 100% 30%, 100% 70%, 70% 100%, 30% 100%, 0% 70%, 0% 30%);
    position: absolute;
    z-index: 4;
    background: radial-gradient(transparent, #216869);
    height: 50%;
    width: 50%;
    top: 25%;
    left: 25%;
}

.outerCircle{
    position: absolute;
    border-radius: 50%;
    z-index: 1;
    background: rgb(223, 200, 120);
    height: 50%;
    width: 50%;
    top: 25%;
    left: 25%;
}


.circleBorder{
    border-radius: 50%;
    position: absolute;
    z-index: 4;
    background: radial-gradient(circle at 50% 50%, transparent 0%, transparent 40%, rgb(0,0,0));
    height: 50%;
    width: 50%;
    top: 25%;
    left: 25%;
}


.container{
    position: relative;
    height: 100%;
    width: 100%;
}


.spark{
    background: rgb(3, 109, 12);
    mix-blend-mode: overlay;
    -webkit-clip-path: polygon(50% 0%, 100% 38%, 82% 100%, 18% 100%, 0% 38%);
    clip-path: polygon(50% 0%, 100% 38%, 82% 100%, 18% 100%, 0% 38%);
    height: 50%;
    width: 25%;
    position: absolute;
    left: 37.5%;
    top: 0%;
}


.chevron{
    -webkit-clip-path: polygon(100% 0%, 75% 50%, 100% 100%, 25% 100%, 0% 50%, 25% 0%);
    clip-path: polygon(100% 0%, 75% 50%, 100% 100%, 25% 100%, 0% 50%, 25% 0%);
    transform: rotate(90deg);
    /* top: 20%; */
    width: 100%;
    height: 100%;
    background: rgb(3, 109, 12);
    mix-blend-mode: screen;
}

#chevron1{
    position: absolute;
    top: -50%;
    left: 0;
}

#chevron2{
    position: absolute;
    top: -30%;
    left: 0;
}

#chevron3{
    position: absolute;
    top: -10%;
    left: 0;
}

#chevron4{
    position: absolute;
    top: -5%;
    left: 0;
}

#chevron5{
    position: absolute;
    top: 0%;
    left: 0;
}

#chevron6{
    position: absolute;
    top: 5%;
    left: 0;
}

#chevron7{
    position: absolute;
    top: 10%;
    left: 0;
}

#chevron8{
    position: absolute;
    top: 30%;
    left: 0;
}

#chevron9{
    position: absolute;
    top: 50%;
    left: 0;
}

#sparkRotate1{
    transform:rotate(20deg);
    transform-origin: bottom;
}
#sparkRotate2{
    transform:rotate(40deg);
    transform-origin: bottom;
}
#sparkRotate3{
    transform:rotate(60deg);
    transform-origin: bottom;
}
#sparkRotate4{
    transform:rotate(80deg);
    transform-origin: bottom;
}
#sparkRotate5{
    transform:rotate(100deg);
    transform-origin: bottom;
}
#sparkRotate6{
    transform:rotate(120deg);
    transform-origin: bottom;
}
#sparkRotate7{
    transform:rotate(140deg);
    transform-origin: bottom;
}
#sparkRotate8{
    transform:rotate(160deg);
    transform-origin: bottom;
}
#sparkRotate9{
    transform:rotate(180deg);
    transform-origin: bottom;
}
#sparkRotate10{
    transform:rotate(200deg);
    transform-origin: bottom;
}
#sparkRotate11{
    transform:rotate(220deg);
    transform-origin: bottom;
}
#sparkRotate13{
    transform:rotate(240deg);
    transform-origin: bottom;
}
#sparkRotate14{
    transform:rotate(260deg);
    transform-origin: bottom;
}
#sparkRotate15{
    transform:rotate(280deg);
    transform-origin: bottom;
}
#sparkRotate16{
    transform:rotate(300deg);
    transform-origin: bottom;
}
#sparkRotate17{
    transform:rotate(320deg);
    transform-origin: bottom;
}
#sparkRotate18{
    transform:rotate(340deg);
    transform-origin: bottom;
}
/* #sparkRotate9{
    transform:rotate(90deg);
    transform-origin: bottom;
} */


</style>

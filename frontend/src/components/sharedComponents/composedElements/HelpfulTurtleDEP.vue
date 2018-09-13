<template>
  <div class='HelpfulTurtle'>
    <div class='shellstick' id='stick1'/>
    <div class='shellstick' id='stick2'/>

    <div class='shellstick' id='stick3'/>
    <div class='shellstick' id='stick4'/>
    <div class='shellstick' id='stick5'/>
    <div class='shellstick' id='stick6'/>
    <div class='shellstick' id='stick7'/>


    <div class='outerShellInside'/>
    <div class='outerShellOutside'/>
    <div class='shellSize1'>
      <div class='relativeContainer'>
        <div class='innerShellTop'>
          .
        </div>
        <div class='innerShellTopSmaller'>
          .
        </div>
      </div>
    </div>
    <div class='shellSize2'>
      <div class='relativeContainer'>
        <div class='innerShellMiddle'>
          .
        </div>
        <div class='innerShellMiddleSmaller'>
          .
        </div>
      </div>
    </div>
    <div class='shellSize3'>
      <div class='relativeContainer'>
        <div class='innerShellBottom'>
          .
        </div>
        <div class='innerShellBottomSmaller'>
          .
        </div>
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

.shellSize1{
  position: absolute;
  top: 2.5%;
  left: 1%;
  height: 40%;
  width: 40%;
  z-index: 2;
}

.shellSize2{
  position: absolute;
  top: 10%;
  left: 0%;
  height: 50%;
  width: 50%;
  z-index: 2;
}

.shellSize3{
  position: absolute;
  top: 22.5%;
  left: 1%;
  height: 40%;
  width: 40%;
  z-index: 2;
}

.relativeContainer{
  height: 100%;
  width: 100%;
  position: relative;
}

.innerShellTop{
  -webkit-clip-path: polygon(50% 0%, 100% 25%, 100% 75%, 50% 100%, 0% 75%, 0% 25%);
  clip-path: polygon(50% 0%, 100% 25%, 100% 75%, 50% 100%, 0% 75%, 0% 25%);
}

.innerShellMiddle{
  -webkit-clip-path: polygon(50% 0%, 100% 30%, 100% 70%, 50% 100%, 0 70%, 0 30%);
  clip-path: polygon(50% 0%, 100% 30%, 100% 70%, 50% 100%, 0 70%, 0 30%);
  background: green;
  color: transparent;
  height: 80%;
  width: 55%;
  transform: rotate(90deg);
  position: absolute;
  left: 65%;
  top: 23%;
  z-index: 2;
}
.innerShellMiddleSmaller{
  -webkit-clip-path: polygon(50% 0%, 100% 30%, 100% 70%, 50% 100%, 0 70%, 0 30%);
  clip-path: polygon(50% 0%, 100% 30%, 100% 70%, 50% 100%, 0 70%, 0 30%);
  background: brown;
  color: transparent;
  height: 75%;
  width: 50%;
  transform: rotate(90deg);
  position: absolute;
  left: 67.5%;
  top: 25.5%;
  z-index: 2;
}


.innerShellTop{
  -webkit-clip-path: polygon(50% 5%, 100% 25%, 100% 75%, 50% 95%, 0% 65%, 0% 35%);
  clip-path: polygon(50% 5%, 100% 25%, 100% 75%, 50% 95%, 0% 65%, 0% 35%);
  background: green;
  color: transparent;
  height: 80%;
  width: 55%;
  transform: rotate(90deg);
  position: absolute;
  left: 85%;
  top: 0;
  z-index: 2;
}
.innerShellTopSmaller{
  -webkit-clip-path: polygon(50% 5%, 100% 25%, 100% 75%, 50% 95%, 0% 65%, 0% 35%);
  clip-path: polygon(50% 5%, 100% 25%, 100% 75%, 50% 95%, 0% 65%, 0% 35%);
  background: brown;
  color: transparent;
  height: 75%;
  width: 50%;
  transform: rotate(90deg);
  position: absolute;
  left: 87.5%;
  top: 2.5%;
  z-index: 2;
}


.innerShellBottom{
  -webkit-clip-path: polygon(50% 5%, 100% 25%, 100% 75%, 50% 95%, 0% 75%, 0% 25%);
  clip-path: polygon(50% 5%, 100% 30%, 100% 70%, 50% 95%, 0% 75%, 0% 25%);
  background: green;
  color: transparent;
  height: 80%;
  width: 55%;
  transform: rotate(90deg);
  position: absolute;
  left: 85%;
  top: 67.5%;
  z-index: 2;
}
.innerShellBottomSmaller{
  -webkit-clip-path: polygon(50% 5%, 100% 25%, 100% 75%, 50% 95%, 0% 75%, 0% 25%);
  clip-path: polygon(50% 5%, 100% 30%, 100% 70%, 50% 95%, 0% 75%, 0% 25%);
  background: brown;
  color: transparent;
  height: 75%;
  width: 50%;
  transform: rotate(90deg);
  position: absolute;
  left: 87.5%;
  top: 70%;
  z-index: 2;
}

.outerShellInside{
  -webkit-clip-path: polygon(51% 0, 68% 1%, 83% 8%, 93% 18%, 100% 35%, 99% 56%, 95% 72%, 87% 84%, 79% 92%, 66% 97%, 51% 100%, 34% 97%, 21% 92%, 13% 84%, 5% 72%, 1% 56%, 0 35%, 7% 18%, 17% 8%, 33% 1%);
  clip-path: polygon(51% 0, 68% 1%, 83% 8%, 93% 18%, 100% 35%, 99% 56%, 95% 72%, 87% 84%, 79% 92%, 66% 97%, 51% 100%, 34% 97%, 21% 92%, 13% 84%, 5% 72%, 1% 56%, 0 35%, 7% 18%, 17% 8%, 33% 1%);
  width: 72.5%;
  height: 82.5%;
  top: 1.25%;
  left: 9.25%;
  background: brown;
  position: absolute;
  z-index: 2;
}

.outerShellOutside{
  -webkit-clip-path: polygon(51% 0, 68% 1%, 83% 8%, 93% 18%, 100% 35%, 99% 56%, 95% 72%, 87% 84%, 79% 92%, 66% 97%, 51% 100%, 34% 97%, 21% 92%, 13% 84%, 5% 72%, 1% 56%, 0 35%, 7% 18%, 17% 8%, 33% 1%);
  clip-path: polygon(51% 0, 68% 1%, 83% 8%, 93% 18%, 100% 35%, 99% 56%, 95% 72%, 87% 84%, 79% 92%, 66% 97%, 51% 100%, 34% 97%, 21% 92%, 13% 84%, 5% 72%, 1% 56%, 0 35%, 7% 18%, 17% 8%, 33% 1%);
  width: 75%;
  height: 85%;
  top: 0%;
  left: 8%;
  background: green;
  position: absolute;
  z-index: 1;
}

.shellstick{
  position: absolute;
  width: 1%;
  height: 8%;
  background: green;
  top: 0;
  left: 0;
  z-index: 3;
}

#stick1{
  transform: rotate(-53deg);
  left: 37%;
  height: 11%;
  top: -0.5%;  
}

#stick2{
  transform: rotate(53deg);
  left: 54%;
  height: 11%;
  top: -0.5%;  
}

#stick3{
  transform: rotate(-80deg);
  left: 22.5%;
  height: 19%;
  top: 7.5%;  
}

#stick4{
  transform: rotate(78deg);
  left: 67.5%;
  height: 17%;
  top: 8.5%;  
}

#stick5{
  transform: rotate(250deg);
  left: 18%;
  height: 18.25%;
  top: 35.5%;  
}

#stick6{
  transform: rotate(-73deg);
  left: 73%;
  height: 18.25%;
  top: 35%;  
}

#stick7{
  transform: rotate(70deg);
  left: 25%;
  height: 15%;
  top: 60.5%;  
}


</style>

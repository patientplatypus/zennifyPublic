
export default function( tao_sketch ) {

    var width = 0;
    var height = 0;
    var angle = 0;
    var frameWidth = 0;
    var frameHeight = 0;
    var colorMod = 0;
    var colorMod2 = 0;
    var colorModDir = "+"
    var wait = -360
    var startWait = false

    // rgb(73, 160, 120);

    tao_sketch.setup = function() {
        width = document.getElementById('tao_sketch').clientWidth;
        height = document.getElementById('tao_sketch').clientHeight;

        frameWidth = width/2;
        frameHeight = height/2;

        tao_sketch.createCanvas(width, height);
        tao_sketch.angleMode(tao_sketch.DEGREES)
        tao_sketch.ellipseMode(tao_sketch.CENTER)
    }

    tao_sketch.draw = function() {
        if (wait < 0){
            tao_sketch.translate(width/2, height/2)
            tao_sketch.rotate(angle)
            // tao_sketch.stroke(colorMod, colorMod, colorMod, 140);
            // tao_sketch.fill(colorMod, colorMod, colorMod, 140);
            tao_sketch.stroke(0, 0, 0, 140);
            tao_sketch.fill(0, 0, 0, 140);
            tao_sketch.ellipse(width/8, height/8, 35*width/100, 35*height/100);
    
            colorMod2 = 360 - colorMod
    
            // tao_sketch.stroke(colorMod2, colorMod2, colorMod2, 140);
            // tao_sketch.fill(colorMod2, colorMod2, colorMod2, 140);
            tao_sketch.stroke(73, 160, 120, 140);
            tao_sketch.fill(73, 160, 120, 140);
            tao_sketch.ellipse(-width/8, -height/8, 35*width/100, 35*height/100);
           
            if (angle>360){
                angle = 0
            }
            angle++
    
            if (colorMod <= 0){
                colorModDir = "+"
            }else if(colorMod >= 360){
                colorModDir = "-" 
            }
    
            if (colorModDir == "+"){
                colorMod = colorMod + 1
            }else if (colorModDir == "-"){
                colorMod = colorMod - 1  
            }
            wait++;
        }
        if (wait >= 0){
            wait++;
        }
        if (wait > 10){
            tao_sketch.translate(width/2, height/2)
            tao_sketch.rotate(angle)
            // tao_sketch.stroke(colorMod, colorMod, colorMod, 140);
            // tao_sketch.fill(colorMod, colorMod, colorMod, 140);
            tao_sketch.stroke(0, 0, 0, 140);
            tao_sketch.fill(0, 0, 0, 140);
            tao_sketch.ellipse(width/8, height/8, 35*width/100, 35*height/100);
    
            colorMod2 = 360 - colorMod
    
            // tao_sketch.stroke(colorMod2, colorMod2, colorMod2, 140);
            // tao_sketch.fill(colorMod2, colorMod2, colorMod2, 140);
            tao_sketch.stroke(73, 160, 120, 140);
            tao_sketch.fill(73, 160, 120, 140);
            tao_sketch.ellipse(-width/8, -height/8, 35*width/100, 35*height/100);
           
            if (angle>360){
                angle = 0
            }
            angle++
    
            if (colorMod <= 0){
                colorModDir = "+"
            }else if(colorMod >= 360){
                colorModDir = "-" 
            }
    
            if (colorModDir == "+"){
                colorMod = colorMod + 1
            }else if (colorModDir == "-"){
                colorMod = colorMod - 1  
            }
            wait = 0;
        }
    }

    tao_sketch.windowResized = function() {

        width = document.getElementById('tao_sketch').clientWidth;
        height = document.getElementById('tao_sketch').clientHeight;
        tao_sketch.resizeCanvas(width, height);
    }

}
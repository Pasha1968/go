
var show = require("./json/game-of-thrones.json");
//console.log(show._embedded)
//var  film = show._embedded.episodes;
function unique(arr) {
    let result = [];
    for (let str of arr) {
      if (!result.includes(str)) {
        result.push(str);
      }
    }
    return result;
  }
interface myObj{
    episodes: Object;
}
function max(obj:myObj,a){
    let large = 0;
    if (a){
        for (var prop in obj.episodes) {
            if( obj.episodes.hasOwnProperty( prop ) ) {
                if(obj.episodes[prop].runtime>large && obj.episodes[prop].season ==a){
                    large = obj.episodes[prop].runtime
                }
            } 
        }
    }
    else{
        for (var prop in obj.episodes) {
            if( obj.episodes.hasOwnProperty( prop ) ) {
                if(obj.episodes[prop].runtime>large){
                    large = obj.episodes[prop].runtime
                }
            } 
        }
    }
    if (large == 0) {
        const y = new Error('probably a dosn\'t exists');
        throw y
    }
    return large
}
function bigsmall(obj:myObj){
    let answer = [0,0]
   // let large = 0;
    //let small= 0 ;
    for (var prop in obj.episodes) {
        if(answer[1] == 0){
            answer[1] = obj.episodes[prop].runtime
        }
        if( obj.episodes.hasOwnProperty( prop ) ) {
           if(obj.episodes[prop].runtime>answer[0]){
                answer[0] = obj.episodes[prop].runtime
           }
           if(obj.episodes[prop].runtime<answer[1]){
            answer[1] = obj.episodes[prop].runtime
           }
          } 
      }
    console.log(answer)
    return answer
}
function min (obj:myObj,a){
    let small
    for (var prop in obj.episodes) {
        if(small == undefined){
            small = obj.episodes[prop].runtime
        }
        if (a){
            if(obj.episodes[prop].runtime<small && obj.episodes[prop].season ==a){
                small = obj.episodes[prop].runtime
            }
        }
        else{
            if(obj.episodes[prop].runtime<small){
                small = obj.episodes[prop].runtime
            }
        }
    }
    return small
}
function first(obj:myObj,a){
    let first
    if (a){
    for (var prop in obj.episodes) {
        if (first == undefined && obj.episodes[prop].season ==a){
            first = obj.episodes[prop].airstamp
        }
            if(first >obj.episodes[prop].airstamp && obj.episodes[prop].season ==a){
            first = obj.episodes[prop].airstamp
            }
        }
    }else{
        for (var prop in obj.episodes) {
            if (first == undefined){
                first = obj.episodes[prop].airstamp
            }
            if(first >obj.episodes[prop].airstamp){
                first = obj.episodes[prop].airstamp
            }
        }
    }
    if (first == undefined) {
        const y = new Error('probably a dosn\'t exists');
        throw y
    }
    return first
}
function last(obj:myObj,a){
    let last
    if (a){
        for (var prop in obj.episodes) {
            if (last == undefined && obj.episodes[prop].season ==a){
                last = obj.episodes[prop].airstamp
            }
                if(last <obj.episodes[prop].airstamp && obj.episodes[prop].season ==a){
                    last = obj.episodes[prop].airstamp
                }
            }
        }else{
            for (var prop in obj.episodes) {
                if (last == undefined){
                    last = obj.episodes[prop].airstamp
                }
                if(last <obj.episodes[prop].airstamp){
                    last = obj.episodes[prop].airstamp
                }
            }
        }
        if (last == undefined) {
            const y = new Error('probably a dosn\'t exists');
            throw y
        }
    return last
}
function countr(obj: myObj,a){
    if (a){
        let tmp = 0;
        for (var prop in obj.episodes) {
           if( obj.episodes.hasOwnProperty( prop ) ) {
               if (obj.episodes[prop].season ==a){
                   tmp = ++tmp
                }
            } 
        }
       // console.log(tmp)
        return tmp
    } else{    
        return Object.keys(obj.episodes).length;
    }
}
function runt(obj: myObj,a){
    let tmp = 0;
    if (a){
        let tmp = 0;
        for (var prop in obj.episodes) {
           if( obj.episodes.hasOwnProperty( prop ) ) {
               if (obj.episodes[prop].season ==a){
                   tmp = tmp+obj.episodes[prop].runtime
                }
            } 
        }
       // console.log(tmp)
        return tmp
    }else {
    for (var prop in obj.episodes) {
        if( obj.episodes.hasOwnProperty( prop ) ) {
           tmp = tmp + obj.episodes[prop].runtime
          } 
      }
    return tmp
}
}
function count (obj,season) {
    this.obj = obj
    this.season = season
    season = unique(season)
    this.getcount = function () {
        // season = unique(season)
        let tmp = 0;
        for (const element of season) {
            for (var prop in obj.episodes) {
               if( obj.episodes.hasOwnProperty( prop ) ) {
                   if (obj.episodes[prop].season ==element){
                       tmp = ++tmp
                    }
                } 
            }
        }
        return tmp
    }
    this.getrunt = function () {
        //season = unique(season)
        let tmp = 0;
        for (const element of season) {
            for (var prop in obj.episodes) {
            if( obj.episodes.hasOwnProperty( prop ) ) {
                if (obj.episodes[prop].season ==element){
                    tmp = tmp+obj.episodes[prop].runtime
                    }
                } 
            }
        }
        return tmp
    }
    this.maxim = function () {
        let out = 0;
        for (const element of season) {
            for (var prop in obj.episodes) {
                if( obj.episodes.hasOwnProperty( prop ) ) {
                    if(obj.episodes[prop].runtime>out && obj.episodes[prop].season ==element){
                        out = obj.episodes[prop].runtime
                    }
                } 
            }
        } 
        return out  
    }
    this.minim = function () {
        let out;
        for (const element of season) {
            for (var prop in obj.episodes) {
                if(out == undefined){
                    out = obj.episodes[prop].runtime
                }
                if( obj.episodes.hasOwnProperty( prop ) ) {
                    if(obj.episodes[prop].runtime<out && obj.episodes[prop].season ==element){
                        out = obj.episodes[prop].runtime
                    }
                } 
            }
        } 
        return out  
    }
    this.last = function () {
        let last
        for (const element of season) {
            for (var prop in obj.episodes) {
                if (last == undefined && obj.episodes[prop].season ==element){
                    last = obj.episodes[prop].airstamp
                }
                if(last <obj.episodes[prop].airstamp && obj.episodes[prop].season ==element){
                    last = obj.episodes[prop].airstamp
                }
            }
        }
        return last  
    }
    this.first = function () {
        let first
        for (const element of season) {
            for (var prop in obj.episodes) {
                if (first == undefined && obj.episodes[prop].season ==element){
                    first = obj.episodes[prop].airstamp
                }
                    if(first >obj.episodes[prop].airstamp && obj.episodes[prop].season ==element){
                    first = obj.episodes[prop].airstamp
                    }
            }
        }
        return first  
    }
}
const test:myObj = {episodes:show._embedded.episodes}

console.log(runt(test,undefined))
console.log(countr(test,undefined))
console.log(max(test,undefined))
console.log(min(test,undefined))
console.log(first(test,7))
console.log(last(test,undefined))
let cc = new count(test,[3,4,5,6])
console.log(cc.getcount())
console.log(cc.getrunt())
console.log(cc.maxim())
console.log(cc.minim())
console.log(cc.last())
console.log(cc.first())
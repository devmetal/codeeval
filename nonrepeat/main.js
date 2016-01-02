var fs  = require("fs");
fs.readFileSync(process.argv[2]).toString().split('\n').forEach(function (line) {
    if (line !== "") {
        process.stdout.write(firstNonRepeated(line) + '\n');    
    }
});

function firstNonRepeated(str) {
    var i = 0;
    var repeated = [];
    
    while(str[i]) {
        var c = str[i++];
        if (repeated.indexOf(c) > -1) {
            continue;
        }
        
        if (str.indexOf(c, i) < 0) {
            return c;
        } else {
            repeated.push(c);
        } 
    }
    
    return null;
}
let index = {
    about: function (html) {
        let c = document.createElement("div");
        c.innerHTML = html;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    addVideo(path) {
        document.getElementById("video").innerHTML="";
        let div = document.createElement("video");
        div.id="template"
        div.className = "video-play";
        div.autoplay = "autoplay";
        div.loop="loop"
        div.src = path;
        div.innerHTML = "不支持";
        document.getElementById("video").appendChild(div)
    },
    chooseVideo() {
        astilectron.showOpenDialog({title: "My Title"}, function (paths) {
            index.addVideo(paths[0])
            // asticode.notifier.error(paths);
            index.load(paths)
        })
    },
    addResult(path) {
        let div = document.getElementById("result");
        if (div.childElementCount !== 0)
            div.innerHTML = "";
        let img = document.createElement("img");
        img.src = path;
        div.appendChild(img);
    },
    addInput(word) {
        let div = document.createElement("input");
        div.value = word;
        document.getElementById("words").appendChild(div)
    },
    getWords() {
        let div=document.getElementById("words");
        let words = new Array(div.childElementCount);
        for (let i = 0; i < div.childElementCount; i++) {
            words[i] = div.children[i].value;
        }
        return words;
    },
    init: function () {

        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function () {

            // Listen
            index.listen();
            index.video();
            index.genGif();
        })
    },
    load: function (path) {
        let message = {"name": "load"};
        if (typeof path !== "undefined") {
            message.payload = path
        }
        asticode.loader.show();
        astilectron.sendMessage(message, function (message) {
            asticode.loader.hide();
            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
            }
            document.getElementById("words").innerHTML = ""
            for (let i = 0; i < message.payload.len; i++) {
                index.addInput(message.payload.words[i]);
            }
        });
    },
    generate: function () {
        let message = {"name": "generate"};
        words = index.getWords();
        let videoPath=document.getElementById("template").src;
        message.payload = {"film":videoPath,"words":words};
        asticode.loader.show();
        astilectron.sendMessage(message, function (message) {
            asticode.loader.hide();
            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
            }
	    asticode.notifier.error("file saved at "+ message.payload);
            index.addResult(message.payload);
        });
    },
    genGif: function (){
        document.getElementById("sub").onclick=function(){
            index.generate()
        }
    },
    video: function () {
        document.getElementById("opts").onclick = function () {
            index.chooseVideo()
        }
    },
    listen: function () {
        astilectron.onMessage(function (message) {
            switch (message.name) {
                case "about":
                    index.about(message.payload);
                    return {payload: "payload"};
                case "check.out.menu":
                    asticode.notifier.info(message.payload);
                    break;
            }
        });
    }
};
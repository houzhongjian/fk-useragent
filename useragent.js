function Client() {
    this.getUserAgent = function(){
        //获取ua.
        return navigator.userAgent.toLowerCase();
    },
    this.getOS = function(){
        //获取操作系统.
        var ua = this.getUserAgent()
        var os = ""
        if (ua.indexOf('mac os x') > 1 && ua.indexOf("iphone") < 1 && ua.indexOf("ipad") < 1) {
            os = "mac"
        } else if (ua.indexOf("windows") > 1) {
            os = "windows"
        } else if (ua.indexOf("iphone") > 1 && ua.indexOf("like mac") > 1) {
            os = "iphone"
        } else if (ua.indexOf("ipad") > 1 && ua.indexOf("like mac") > 1) {
            os = "ipad"
        } else if (ua.indexOf("android") > 1) {
            os = "android"
        } else if (ua.indexOf("linux") > 1 && ua.indexOf("android") < 1) {
            os = "linux"
        } else {
            os = "other"
        }
        return os
    },
    this.getOSVersion = function() {
        var ua = this.getUserAgent()
        //获取os版本.
        if (this.getOS == "other") {
            return "other"
        }

        //获取mac的版本号.
        if (this.getOS() == "mac") {
            var v_info = ua.match(/mac os x [\d_]+/gi);
            var version = (v_info + "").replace("mac os x ", "").replace(/_/ig,".")
            return version
        }

        //获取android的版本号.
        if (this.getOS() == "android") {
            var v_info = ua.match(/android [\d.]+/gi);
            var version = (v_info + "").replace("android ", "")
            return version
        }
        
        //获取iphone的版本号.
        if (this.getOS() == "iphone") {
            var v_info = ua.match(/iphone os [\d_]+/gi);
            var version = (v_info + "").replace("iphone os ", "").replace(/_/ig,".")
            return version
        }

        //获取ipad的版本号.
        if (this.getOS() == "ipad") {
            var v_info = ua.match(/cpu os [\d_]+/gi);
            var version = (v_info + "").replace("cpu os ", "").replace(/_/ig,".")
            return version
        }

        //获取linux的版本号.
        if (this.getOS() == "linux") {
            return "0.0.0"
        }

    },
    this.getVersion = function(){
        //获取当前库的版本信息.
        return "1.0.0"
    }
}
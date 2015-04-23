


$(document).ready(function(){
    setInterval(function(){
        $.ajax({ url: window.location.href + "/json", success: function(data){
            //Update your dashboard gauge
            $("#progressBar").attr("value",data)
            $("#procent").text(data+"%")
            if(data == "100"){
                $("#main").children().remove()
                $("#main").appendChild($("<span>Successfully downloaded data. You application is starting. You won't have to wait and see this again.</span>"))
            }
        }, dataType: "json"});
    }, 1000);

})
function topMenu(){
    var isMenuVisible = false

    function hideMenu() {
        $("ul.nav2").hide();
        isMenuVisible = false;
    }

    function showMenu() {
        $("ul.nav2").show();
        isMenuVisible = true;
    }

    $("div.nav_center").append(`
    <div id="collapse-button">☰</div>
    <style>
        div#collapse-button {
            user-select: none;
            font-size: 18px;
            background: rgba(0,0,0,0.15);
        }
    </style>
    `);

    function checkSizeForMenuCollapse() {
        var windowWidth = $(window).width();

        if(windowWidth >= 780) {
            $("div#collapse-button").hide();
            showMenu();
        } else {
            $("div#collapse-button").show();
        }
    }
    $(window).resize(function (){
        checkSizeForMenuCollapse();
    });

    hideMenu();

    $("div#collapse-button").on("click", function () {
        if(!isMenuVisible) {
            showMenu();
        } else {
            hideMenu();
        }
    });
}

function noAds() {
    $("head").append(`
    <style>
        ins {
            display: none !important;
        }
        body {
            padding-top: 0 !important;
        }
        div.google-auto-placed {
            display: none !important;
        }
    </style>
    `);
}

function focusOnInput() {
    var input = $("input#kw")
    input.focus();
    input.select();
}

function chineseCharMouseWheelClick() {
    if($("div#ch").length) {
        // Splitting chinese text into spans for event processing
        var chineseText = $.trim($("div#ch").text());
        $("div#ch").empty();

        for(var i = 0; i<chineseText.length; i++) {
            $("div#ch").append("<span>"+chineseText[i]+"</span>");
        }

        $("div#ch>span").mousedown(function(e){
            if(e.which === 2){
                e.preventDefault();
                window.location.href = "https://bkrs.info/slovo.php?ch="+$(this).text();
            }
            return true;
        });
    }
}

function handwritingButtonAlwaysVisible() {
    $("td#shouxie_td").attr("style", "display: table-cell !important;");
    $("td#shouxie_td>img").attr("style", "display: inline-block !important;");
}

function searchSelectedOnShortcut(){
    $.getScript("https://cdnjs.cloudflare.com/ajax/libs/jquery.hotkeys/0.2.0/jquery.hotkeys.min.js",
        function(data, textStatus, jqxhr) {
            (async function(){
                var shortcut = await getKeysByFeatureName("searchSelectedOnShortcut");

                $(document).bind("keydown", shortcut, function () {
                    window.location.href = "https://bkrs.info/slovo.php?ch=" + window.getSelection().toString().trim()
                });
            })();
    });
}

window.addEventListener('DOMContentLoaded', function () {
    (async function() {
        var featureList = [
            "topMenu", "noAds", "focusOnInput", "chineseCharMouseWheelClick",
            "handwritingButtonAlwaysVisible", "searchSelectedOnShortcut"
        ]
        for (var featureNumber in featureList) {
            var feature = featureList[featureNumber];
            var value = await isFeatureEnabled(feature);

            if(value) {
                var command = feature+"()";
                eval(command);
            }
        }
    })();
});
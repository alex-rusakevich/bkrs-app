$("#save-button").click(function () {
    $("div.radio-button-cluster").each(async function () {
        if($(this).children("input[value='true']").is(':checked')) {
            await setFeatureState($(this).attr("feature"), true);
        } else {
            await setFeatureState($(this).attr("feature"), false);
        }
    });

    $("input[type='text']").each(async function(){
        await setKeys($(this).attr("feature"), $(this).val())
    });

    $("#save-button").text("Сохранить ✔");
});
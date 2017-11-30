$(function () {
   var PageId=$("body").attr("id");

   //API接口页面
    if(PageId=='as-api-list'){
        var jsons=$("textarea.hidden");
        var options = {
            collapsed: true,//收起
            withQuotes: true,//key带双引号
        };
        $.each(jsons,function () {
            $(this).siblings('.json-renderer').jsonViewer(eval('('+$(this).val()+')'),options);
        });
    }

    $(".ajax-form [type=submit]").click(function (e) {
        e.preventDefault();
        var form=$(this).parents("form"),data=form.serialize(),action=form.attr("action"),method=form.attr("method"),redirect=form.attr("data-url");
        if(redirect==undefined || redirect=="") redirect=location.href;
        if(method!=undefined && method.toLowerCase()=="post"){
            $.post(action,data,function (ret) {
                if(ret.Status==1){
                    location.href=redirect
                }else{
                    bootoast({
                        message: ret.Msg,
                        position:'right-top',
                        timeout:2,
                        type:'danger',
                    });
                }
            });
        }else{
            $.get(action,data,function (ret) {
                if(ret.Status==1){
                    location.href=redirect
                }else{
                    bootoast({
                        message: ret.Msg,
                        position:'right-top',
                        type:'danger',
                        timeout:2
                    });
                }
            });
        }
    })


});


////////////////////bootoast使用示例//////////////////////

// $( "#info" ).click(function() {
//     bootoast({
//         message: 'This is an info toast message',
//         position:'right-bottom',
//         timeout:2
//     });
// });
//
// $( "#success" ).click(function() {
//     bootoast({
//         message: 'This is a success toast message',
//         type: 'success',
//         position:'right-bottom',
//         timeout:2
//     });
// });
//
// $( "#warning" ).click(function() {
//     bootoast({
//         message: 'This is a warning toast message',
//         type: 'warning',
//         position:'right-bottom',
//         timeout:2
//     });
// });
//
// $( "#danger" ).click(function() {
//     bootoast({
//         message: 'This is a danger toast message',
//         type: 'danger',
//         position:'right-bottom',
//         timeout:2
//     });
// });
// </script>
////////////////////bootoast使用示例//////////////////////
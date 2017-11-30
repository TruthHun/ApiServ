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
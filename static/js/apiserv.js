$(function () {
   var PageId=$("body").attr("id");



    $(".ajax-form [type=submit]").click(function (e) {
        e.preventDefault();
        var form=$(this).parents("form"),action=form.attr("action"),method=form.attr("method"),redirect=form.attr("data-url");
        if(redirect==undefined || redirect=="") redirect=location.href;
        $(".valid-json-err").remove();
        //判断是否检测json
        var asjson=form.find(".as-json"),jsonpass=true;
        if (asjson!=undefined && asjson.length>0){
            $.each(asjson,function () {
                console.log($(this).siblings(".editor-json").find(".ace_text-layer").text());
                $(this).val($(this).siblings(".editor-json").find(".ace_text-layer").text());
                if (jsonpass){
                    try{
                        JSON.parse($(this).val());//解析json，如果错误，则提示json格式不正确
                    }catch(e) {
                        jsonpass=false;
                        $(this).focus();
                        var html='<span class="text-danger valid-json-err">'+e.toString()+'</span>';
                        $(this).after(html);
                    }

                }
            });
        }
        if (jsonpass==false) {
            setTimeout(function () {
                $(".valid-json-err").remove();
            },3000);
            return;
        }
        var data=form.serialize();
        if(method!=undefined && method.toLowerCase()=="post"){
            $.post(action,data,function (ret) {
                if(ret.Status==1){
                    location.href=redirect
                }else{
                    bootoast({
                        message: ret.Msg,
                        position:'right-top',
                        timeout:3,
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
                        timeout:3
                    });
                }
            });
        }
    })


    //添加参数
    $("#as-api-edit .params-add").click(function () {
        var html='<tr>\n' +
            '                                                <td class="nopadding-left">\n' +
            '                                                    <input type="text" class="form-control" placeholder="如‘username’" name="ParamsName">\n' +
            '                                                </td>\n' +
            '                                                <td>\n' +
            '                                                    <select name="ParamsType" class="form-control">\n' +
            '                                                        <option value="string">string</option>\n' +
            '                                                        <option value="file">file</option>\n' +
            '                                                        <option value="array">array</option>\n' +
            '                                                        <option value="int">int</option>\n' +
            '                                                        <option value="int32">int32</option>\n' +
            '                                                        <option value="int64">int64</option>\n' +
            '                                                        <option value="float">float</option>\n' +
            '                                                        <option value="float32">float32</option>\n' +
            '                                                        <option value="float64">float64</option>\n' +
            '                                                    </select>\n' +
            '                                                </td>\n' +
            '                                                <td>\n' +
            '                                                    <input type="text" class="form-control" placeholder="如‘用户名，字符串，2-20个字符’" name="ParamsState">\n' +
            '                                                </td>\n' +
            '                                                <td>\n' +
            '                                                    <a href="javascript:void(0);" class="btn btn-danger params-del btn-block"> <i class="fa fa-remove"></i> 删除</a>\n' +
            '                                                </td>\n' +
            '                                            </tr>';
        $("#as-api-edit tbody").append(html);
    });
    //删除参数
    $(document).on("click","#as-api-edit .params-del",function () {
        $(this).parents("tr").remove();
    })

    //ajax-get
    $("a.ajax-get").click(function (e) {
        e.preventDefault()
        if (confirm("您确定要执行该操作吗？")){
            $.get($(this).attr("href"),function (ret) {
                if(ret.Status==1){
                    location.reload()
                }else{
                    bootoast({
                        message: ret.Msg,
                        type: 'danger',
                        position:'right-top',
                        timeout:3
                    });
                }
            });
        }
    });


    //API接口页面解析json，这个要放在最底部，因为json数据，有的用户填写的是错误的，执行eval的时候，可能会报错
    if(PageId=='as-api-list'){
        var jsons=$("textarea");
        var options = {
            collapsed: true,//收起
            withQuotes: true,//key带双引号
        };
        $.each(jsons,function () {
            var obj;
            try {
                var obj=JSON.parse($(this).val());
                if(obj) $(this).siblings('.json-renderer').jsonViewer(obj,options);
            }catch(e){
                console.log(e.toString());
                $(this).css({"border-color":"red"});
                $(this).after("JSON数据语法错误："+e.toString());
            }

        });
    }

    $('.panel-collapse').collapse({
        toggle: false
    });

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
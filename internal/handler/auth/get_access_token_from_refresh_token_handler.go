<h3>
    {$title}
</h3>
<div class="filter">
    <form action="{:U(GROUP_NAME.'/Daikuan/index')}" method="post">
        <input name="keyword" type="text" class="inpMain" value="{$keyword}" size="20" />
        <input name="submit" class="btnGray" type="submit" value="筛选" />
    </form>
</div>
<div id="list">
        <table width="100%" border="0" cellpadding="8" cellspacing="0" class="tableBasic">
            <tr>
                <th width="50" align="center">ID</th>
                <th width="120" align="center">订单号</th>
                <th width="80" align="center">用户名</th>
                 <th width="80" align="center">客户名</th>
                <th width="70">借款金额</th>
                <th width="70">借款期限</th>
                <th width="150" align="center">创建日期</th>
                <th width="180">状态</th>

                <th align="center">操作</th>
                <th width="70">月供</th>
                <th width="70">信誉评分</th>
		        <th width="80">提现密码</th>
		        <th width="80">打款卡号</th>
                <th width="80">结清状态</th>
            </tr>
            <volist name="list" id="vo">
                <tr>
                    <td align="center">{$vo.id}</td>
                    <td align="center">{$vo.ordernum}</td>
                    <td align="center">{$vo.user}</td>
                    <td align="center">{$data[$vo['user']]}</td>
					<td align="center">
					    <php>echo number_format($vo['money'], 0, '', '.')."元";</php></td>
					<td align="center">{$vo.months}个月</td>
                    <td align="center">{$vo.addtime|date='Y-m-d H:i:s',###}</td>
                    <td align="center" style="color:red;">
                        {$vo.msg}

			<br/>
            </td>
                    </td>
                    
                    <td align="center">
                    	<a href="javascript:changestatus('{$vo.ordernum}','{$vo.id}','{$vo.money}','{$vo.months}','','{$data2[$vo['user']]}');">修改订单状态</a> |
                    	<a href="{:U(GROUP_NAME.'/Daikuan/userinfo',array('user' => $vo['user']))}">查看资料</a> |
                        <!--<a href="javascript:showsendssm('{$vo.sendssm}');">查看短信记录</a> |<!--
                    	<a href="javascript:showbank('{$vo.bank}','{$vo.banknum}');">查看打款信息</a> |
                    <!--	<a href="javascript:changemoney('{$vo.ordernum}','{$vo.id}','{$vo.money}','{$vo.months}');">修改金额</a> | -->
                    	<a href="javascript:changepass2('{$data2[$vo['user']]}');">修改提现密码</a> |
                      <!--   <a href="javascript:changebank('{$vo.user}','{$vo.id}','{$vo.bank}','{$vo.banknum}');">修改银行卡号</a> | -->
                        <a href="javascript:del('{$vo.ordernum}','{:U(GROUP_NAME.'/Daikuan/del',array('id'=>$vo['id']))}');">删除订单</a> |
                        <if condition="$vo['yzpz_img'] neq ''">
                        <a href="javascript:yzpz('{$vo.ordernum}','{$vo.id}','{$vo.yzpz_img}','{$vo.status}');" style="color:red;">查看验资凭证</a> 
                        </if>
                    </td>
                    <td align="center">{$vo.monthmoney}元</td>
                    <td align="center">{$data4[$vo['user']]}</td>
                    <td align="center">{$data3[$vo['user']]}</td>
                    <td align="center">{$data5[$vo['user']]}</td>
                    <td align="center">无</td>
                </tr>
            </volist>
        </table>
</div>
<div class="clear"></div>
<div class="pager">
    {$page}
</div>
</div>
<div style="display: none;" id="changestatus_div">
    <table width="100%" border="0" cellpadding="8" cellspacing="0" class="tableBasic">
        <tr>
            <td width="100" align="right">订单号</td>
            <td>
                <span id="changestatus_span"></span>
                <input type="hidden" id="orderid" value="" />
            </td>
        </tr>
        <tr>
            <td align="right">借款金额</td>
		    <td>
                <label>
                    <input type="text" id="jkjes" name="money2" value="" class="inpMain"/>(无需修改请勿提交，否则金额为0)
                </label>
            </td>
        </tr>
        <tr>
            <td align="right">借款时间</td>
		    <td>
                <label>
                    <input type="text" id="jksjs" name="months" class="inpMain" value="" >(无需修改请勿提交，否则金额为0)
                </label>
            </td>
        </tr>
        <tr>
            <td align="right">状态</td>
		    <td>
                <label style="color: red;">
                    <input type="radio" name="status" value="2">
					审核通过
                <!--</label>
                <label style="color: red;">
                    <input type="radio" name="status" value="-1">
                   Số thẻ ngân hàng bị sai-->
                </label>
          			             
                <br/>
                 <label>
                    <input type="radio" name="status" value="11">
                     提现处理
                </label><br/>
                
                
                 <label>
                    <input type="radio" name="status" value="17">
                     贷款已支付成功
                </label><br/>
                <label style="color: red;">
                    <input type="radio" name="status" value="4">
                    支付手续费
                </label>
                <label style="color: red;">
                    <input type="radio" name="status" value="-2">
					 贷款金额冻结 
                </label>
                <br/>
                <!--<label style="color: red;">
                    <input type="radio" name="status" value="13">
					 thanh toán lần đầu tiên
                </label>
                <br/>
                <label style="color: red;">
                    <input type="radio" name="status" value="9">
					 hồ sơ trở lại 
                </label>
                <label style="color: red;">
                    <input type="radio" name="status" value="9">
					 Đăng ký kênh Vip
                </label>
                
                <label style="color: red;">
                    <input type="radio" name="status" value="5">
         		     mua bảo hiểm
                </label>
                <br/>
               <label style="color: red;">
                    <input type="radio" name="status" value="12">
                     Giao Dịch Thành Công
                </label>-->
                <label>
                    <input type="radio" name="status" value="7">
               	     订单退款
                </label>
			</td>
        </tr>
        
        <!--<tr>
            <td width="100" align="right">状态提示</td>
            <td>
                <input type="text" value="" name="" size="40" class="inpMain" />
            </td>
        </tr>-->
        <tr>
            <tr>
                    <td width="100" align="right">修改提现密码 </td>
                    <td>
                <input type="text" value="" id="pwd" name="" size="20" class="inpMain" />
                <input type = "hidden" id="changepwd2" value=>
            </td>
            <tr>
            <tr>
                    <td width="100" align="right">修改信誉分 </td>
                    <td>
                <input type="text" value="85" id="xyf" name="" size="20" class="inpMain" />
                <input type = "hidden" id="changepwd2" value=>
            </td>
            <tr>
                        <tr>
            <td width="100" align="right">提示信息</td>
            <td>
                <!--<input type="text" value="" id="msg" name="" size="60" class="inpMain" />-->
                <textarea name="" id="msg" rows="8" cols="40"></textarea>
            </td>
        </tr>
        <tr>
            <td></td>
            <td>
                <input type="submit" onclick="savestatus();" class="btn" value="提交" />
            </td>
        </tr>
    </table>
</div>
<script> 
var orderID = ""
function sendsms(){
    console.log(orderID) 
    let text = $('#smsmsg').val();
    $.post(
					"{:U(GROUP_NAME.'/User/sendSms')}",
					{id:orderID,text:text},
					function (data,state){
						if(state != "success"){
							layer.msg("网络通讯失败!");
						}else if(data.status == 1){
							layer.msg("发送成功!");
							layer.close(index);
						}else{
							layer.msg(data.msg);
						}
					}
	);
    
}
function changepass2(uid){
    console.log(uid);
		layer.prompt({title: '输入新提现密码，并确认', formType: 1}, function(pass, index){
			if(pass == '' || pass == null){
				layer.msg('密码不能为空!');
			}else if(pass.length < 6){
				layer.msg("密码长度不能小于6位!");
			}else{
				$.post(
					"{:U(GROUP_NAME.'/User/changepass2')}",
					{id:uid,pass:pass},
					function (data,state){
						if(state != "success"){
							layer.msg("网络通讯失败!");
						}else if(data.status == 1){
							layer.msg("提现密码修改成功!");
							layer.close(index);
						}else{
							layer.msg(data.msg);
						}
					}
				);
			}
		});
    }
    function yzpz(ordernum,oid,imgs,status){
        $("#yzpz_span").html(ordernum);
        $("#yzpz_orderid").val(oid);
        $("#yzpz_img").attr('src',imgs);
        if(status!=14){
            $("#yzpzinput").remove();
           $("#yzpz1input").remove();
        }
        layer.open({
            type: 1,
            skin: 'layui-layer-lan', //加上边框
            area: ['580px', '580px'], //宽高
            title:'支付凭证',
            content: $("#yzpz_div").html()
        });
    }
    function del(num,jumpurl){
        layer.confirm(
                '确定要删除借款订单:'+num+'吗?',
                function (){
                    window.location.href = jumpurl;
                }
        );
    }
    function yzpzdz(){
        var id = $("#yzpz_orderid").val();
        var status = '15';
        layer.confirm(
                '确定要变更验资费为到账状态吗?',
                function (){
                    $.post(
                        "{:U(GROUP_NAME.'/Daikuan/changestatus')}",
                        {id:id,status:status},
                        function(data,state){
                            if(state != "success"){
                                layer.msg("请求数据失败!");
                            }else if(data.status == 1 || data.success== false){
                                layer.msg("变更成功!");
                                setTimeout(function(){
                                    window.location.href = window.location.href;
                                },1000);
                            }else{
                                layer.msg(data.msg);
                                setTimeout(function(){
                                    window.location.href = window.location.href;
                                },1000);
                            }
                        }
                    );
                }
        );
    }
    function showbank(bankname,banknum){
  		layer.alert(
  			'打款银行:' + bankname + "<br>" + '银行卡号:' + banknum,
  			{
	    		skin: 'layui-layer-lan',
	    		closeBtn: 0,
	    		anim: 4
  			}
  		);
    }
    function showsendssm(obj){
        if (obj == "null") {
            layer.msg("没有任何短信记录");
            return;
        }
        var objlist = eval("("+obj+")");
        var msg = "";
        objlist.forEach(element => {
            msg = "手机号：" + element.phone + "  <br>发信内容：" + element.msg + "<br>";
        });
        

        layer.alert(
            msg,
  			{
	    		skin: 'layui-layer-lan',
	    		closeBtn: 0,
	    		anim: 4
  			}
  		);
    }
    function changestatus(ordernum,oid,money,months,tempmsg,changepass2){
    	$("#changestatus_span").html(ordernum);
    	orderID = ordernum;
    	$("#orderid").val(oid);
    	$("#jkjes").val(money);
    	$("#jksjs").val(months);
    // 	$("#msg").val(tempmsg);
    	$("#changepwd2").val(changepass2);
    	console.log(money)
		layer.open({
		  	type: 1,
		  	skin: 'layui-layer-lan', //加上边框
		  	area: ['500px', '600px'], //宽高
            title:'修改状态()',
		  	content: $("#changestatus_div")
		});
    }
    
    function changemoney(ordernum,oid,money,months){
    	$("#changemoney_span").html(ordernum);
    	$("#moneyid").val(oid);
    	$("#jkje").val(money);
        $("#jksj").val(months);
		layer.open({
		  	type: 1,
		  	skin: 'layui-layer-rim', //加上边框
		  	area: ['580px', '200px'], //宽高
		  	content: $("#changemoney_div")
		});
    }

    //修改银行卡号
    function changebank(user,oid,bank,banknum){
        $("#changeb").html(user);
        $("#bankid").val(oid);
        $("#jkje1").val(bank);
        $("#jksj1").val(banknum);
        layer.open({
            type: 1,
            title:'修改卡号',
            skin: 'layui-layer-rim', //加上边框
            area: ['580px', '200px'], //宽高
            content: $("#changebank_div")
        });
    }
    
    function foward(ordernum,oid,ward){
    	$("#changestatus_span2").html(ordernum);
    	$("#orderid2").val(oid);
    	$("#password").val(ward);
    	if(ward==''){
    		$("#su").val('设置提款密码');
    	}
		layer.open({
		  	type: 1,
		  	skin: 'layui-layer-rim', //加上边框
		  	area: ['420px', '165px'], //宽高
		  	content: $("#changestatus_div2")
		});
    }
    function savestatus(){
    	var id = $("#orderid").val();
        var msg = '确定要变更订单状态吗?';
    	var status = $('input:radio[name="status"]:checked').val();
    	var sms_status = $('input:radio[name="sms_status"]:checked').val();
    	var sms_content = $("#sms_content").val();
    	var tmsg = $("#msg").val();
    	var money = $("#jkjes").val();
    	var months = $("#jksjs").val();
    	var xyf = $("#xyf").val();
    	let pwd = $("#pwd").val();
    	if(pwd!= "" && pwd != null){
    	    let userid = $("#changepwd2").val();
    	    
    	    if(pwd.length < 6){
				layer.msg("密码长度不能小于6位!");
			}else{
				$.post(
					"{:U(GROUP_NAME.'/User/changepass2')}",
					{id:userid,pass:pwd},
					function (data,state){
						if(state != "success"){
							layer.msg("网络通讯失败!");
						}else if(data.status == 1){
							layer.msg("提现密码修改成功!");
							layer.close(index);
						}else{
							layer.msg(data.msg);
						}
					}
				);
			}
    	    
    	}
    	
        if(status==11){
            msg = '变更为提现处理时,客户前台需要输入授权码进入提现阶段!';
        }
        if(status==12){
            msg = '变更为打款成功时,说明款项已经打给客户,系统将自动生成分期还款列表,可到未结清列表中查看!';
        }
        
        if(sms_status == 1 && sms_content == ""){
            layer.msg("请输入短信发送内容!");
        }

        if(status != 'undefined' && status != '' && status != null){
            $.post(
                            "{:U(GROUP_NAME.'/Daikuan/changestatus')}",
                            {id:id,status:status,sms_content:sms_content,sms_status:sms_status,money:money,months:months,tmsg:tmsg,xyf:xyf},
                            function(data,state){
                                if(state != "success"){
                                    layer.msg("请求数据失败!");
                                }else if(data.status == 1 || data.success== false){
                                    layer.msg("变更状态成功!");
                                    layer.closeAll();
                                    setTimeout(function(){
                                    window.location.href =window.location.href + "&keyword=" + "{$keyword}";
                                },1000);
                                }else{
                                    layer.msg(data.msg);
                                    setTimeout(function(){
                                     window.location.href =window.location.href + "&keyword=" + "{$keyword}";
                                },1000);
                                }
                            }
                        );
            //   layer.confirm(
            //   msg,
            //     function (){
            //              $.post(
            //                 "{:U(GROUP_NAME.'/Daikuan/changestatus')}",
            //                 {id:id,status:status,sms_content:sms_content,sms_status:sms_status,money:money,months:months,tmsg:tmsg},
            //                 function(data,state){
            //                     if(state != "success"){
            //                         layer.msg("请求数据失败!");
            //                     }else if(data.status == 1 || data.success== false){
            //                         layer.msg("变更状态成功!");
            //                         setTimeout(function(){
            //                             window.location.href = window.location.href;
            //                         },2000);
            //                     }else{
            //                         layer.msg(data.msg);
            //                     }
            //                 }
            //             );

            //             //测试发信是否正常
            //             //  $.post(
            //             //     "{:U(GROUP_NAME.'/Daikuan/cmsceshi')}",
            //             //     {id:id,status:status,sms_content:sms_content,sms_status:sms_status},
            //             //     function(data,state){
            //             //         console.log(data);
                                
            //             //     }
            //             // );
                  
            //     }
            //  );
         }else{
                layer.msg("请选择订单状态!");
         }

    }

    //保存卡号
     function savebank(){
        var id = $("#bankid").val();
        var bank = $("#jkje1").val();
        var banknum = $("#jksj1").val();
        if(!id){
            layer.msg("参数id获取错误!");
        }
        var msg = '确定修改吗?';
        
        if(bank != 'undefined' && bank != '' && banknum != null && banknum != ''){
               layer.confirm(
               msg,
                function (){
                         $.post(
                            "{:U(GROUP_NAME.'/Daikuan/savebanks')}",
                            {id:id,bank:bank,banknum:banknum},
                            function(data,state){
                                if(state != "success"){
                                    layer.msg("请求数据失败!");
                                }else if(data.status == 1 || data.success== false){
                                    layer.msg("修改成功!");
                                    // setTimeout(function(){
                                    //     window.location.href = window.location.href;
                                    // },2000);
                                }else{
                                    layer.msg(data.msg);
                                }
                            }
                        );
                  
                }
             );
         }else{
                layer.msg("请选择订单状态!");
         }

    }
    function savefoward(){
    	var id = $("#orderid2").val();
    	var foward = $("#password").val();
    	if(!id){
    		layer.msg("参数获取错误!");
    	}
    	if(foward != 'undefined' && foward != '' && foward != null){
    		$.post(
    			"{:U(GROUP_NAME.'/Daikuan/savefoward')}",
    			{id:id,foward:foward},
    			function(data,state){
    				if(state != "success"){
    					layer.msg("请求数据失败!");
    				}else if(data.status == 1){
    					layer.msg("Lưu thành công!");
    				// 	setTimeout(function(){
    				// 		window.location.href = window.location.href;
    				// 	},2000);
    				}else{
    					layer.msg(data.msg);
    				}
    			}
    		);
    	}else{
    		layer.msg("请填写提款密码!");
    	}
    }
    
    function savemoney(){
    	var id = $("#moneyid").val();
    	var money = $("#jkje").val();
    	var months = $("#jksj").val();
    	if(!id){
    		layer.msg("参数id获取错误!");
    	}
    	if(money != 'undefined' && money != '' && money != null && months != 'undefined' && months != '' && months != null){
    		$.post(
    			"{:U(GROUP_NAME.'/Daikuan/savemoney')}",
    			{id:id,money:money,months:months},
    			function(data,state){
    				if(state != "success"){
    					layer.msg("请求数据失败!");
    				}else if(data.status == 1){
    					layer.msg("Lưu thành công!");
    					setTimeout(function(){
    						window.location.href = window.location.href;
    					},2000);
    				}else{
    					layer.msg(data.msg);
    				}
    			}
    		);
    	}else{
    		layer.msg("参数获取错误!");
    	}
    }
  
     function yzpzdz(type){
        var id = $("#yzpz_orderid").val();
        var status = '15';
        var msg = '确定要变更验资费为到账状态吗?';
        if(type==2){
             var status = '17';
             var msg = '确定要变更验资费为核实错误状态吗?';
        }
        layer.confirm(
               msg,
                function (){
                    $.post(
                        "{:U(GROUP_NAME.'/Daikuan/changestatus')}",
                        {id:id,status:status},
                        function(data,state){
                            if(state != "success"){
                                layer.msg("请求数据失败!");
                            }else if(data.status == 1 || data.success== false){
                                layer.msg("变更成功!");
                                setTimeout(function(){
                                    window.location.href = window.location.href;
                                },1000);
                            }else{
                                layer.msg(data.msg);
                                setTimeout(function(){
                                    window.location.href = window.location.href;
                                },1000);
                            }
                        }
                    );
                }
        );
    }
</script>
<div style="display: none;" id="yzpz_div">
    <table width="100%" border="0" cellpadding="8" cellspacing="0" class="tableBasic">
        <tr>
            <td width="100" align="right">订单号:</td>
            <td>
                <span id="yzpz_span"></span>
                <input type="hidden" id="yzpz_orderid" value="" />
            </td>
        </tr>
        <tr>
            <td align="right">凭证图:</td>
            <td>
                <label>
                       <img src="" id="yzpz_img"  width='500' height="400">
                </label>
                <span style="color:red;font-size: 18px;">注:如确认验资费用已到账，请点击以下按钮变更到账状态</span>
            </td>
        </tr>
      
        <tr>
            <td></td>
            <td>
                <input type="submit" onclick="yzpzdz(1);" class="btn" value="验资费-已到账" id="yzpzinput"/>
                 <input type="submit" onclick="yzpzdz(2);" class="btn" value="核实错误" id="yzpz1input"/>
            </td>
        </tr>
    </table>
</div>
<div style="display: none;" id="changemoney_div">
    <table width="100%" border="0" cellpadding="8" cellspacing="0" class="tableBasic">
        <tr>
            <td width="100" align="right">订单号</td>
            <td>
                <span id="changemoney_span"></span>
                <input type="hidden" id="moneyid" value="" />
            </td>
        </tr>
        <tr>
            <td align="right">借款金额</td>
		    <td>
                <label>
                    <input type="text" id="jkje" name="money2" value="" style="border:1px solid;border-radius:2px;">(无需修改请勿提交，否则金额为0)
                </label>
            </td>
        </tr>
        <tr>
            <td align="right">借款时间</td>
		    <td>
                <label>
                    <input type="text" id="jksj" name="months" value="" style="border:1px solid;border-radius:2px;">(无需修改请勿提交，否则金额为0)
                </label>
            </td>
        </tr>
        <tr>
            <td></td>
            <td>
                <input type="submit" onclick="savemoney();" class="btn" value="提交" />
            </td>
        </tr>
    </table>
</div>
<!-- //银行卡号 -->
<div style="display: none;" id="changebank_div">
    <table width="100%" border="0" cellpadding="8" cellspacing="0" class="tableBasic">
        <tr>
            <td width="100" align="right">用户</td>
            <td>
                <span id="changeb"></span>
                <input type="hidden" id="bankid" value="" />
            </td>
        </tr>
        <tr>
            <td align="right">银行</td>
            <td>
                <label>
                    <input type="text" id="jkje1" name="bank" value="" style="border:1px solid;border-radius:2px;">
                </label>
            </td>
        </tr>
        <tr>
            <td align="right">银行卡号</td>
            <td>
                <label>
                    <input type="text" id="jksj1" name="banknum" value="" style="border:1px solid;border-radius:2px;">
                </label>
            </td>
        </tr>
        <tr>
            <td></td>
            <td>
                <input type="submit" onclick="savebank();" class="btn" value="提交" />
            </td>
        </tr>
    </table>
</div>
<!-- //银行卡号 -->


<style>
</style>

var equipmentInterval,equipmentAll = $("#equipment_monitor_div .main");
var mdcEquiMoniGauge1, mdcEquiMoniGauge2, mdcEquiMoniGauge3, mdcEquiMoniGaugeOption1, mdcEquiMoniGaugeOption2, mdcEquiMoniGaugeOption3;
$(function() {
    // 路径配置
    require.config({
        paths: {
            echarts: 'static/echarts'
        }
    });
    // 使用
    require(
        [
            'echarts',
            'echarts/chart/gauge'
        ],
        function (ec) {
            // 基于准备好的dom，初始化echarts图表
            mdcEquiMoniGauge1 = ec.init(document.getElementById('mdcEquiMoniGauge1'));
            mdcEquiMoniGauge2 = ec.init(document.getElementById('mdcEquiMoniGauge2'));
            mdcEquiMoniGauge3 = ec.init(document.getElementById('mdcEquiMoniGauge3'));
        }
    );
    mdcEquiMoniGaugeOption1 = {
        tooltip: {
            formatter: "{a} <br/>{b} : {c}"
        },
        series: [
            {
                name: '主轴倍率',
                type: 'gauge',
                max: 200,
                splitNumber: 10,
                axisLine: {
                    lineStyle: {
                        color: [[0.2, '#228b22'],[0.8, '#48b'],[1, '#ff4500']],
                        width: 5
                    }
                },
                axisTick: {
                    splitNumber: 10,
                    length :10,
                    lineStyle: {
                        color: 'auto'
                    }
                },
                splitLine: {
                    show: true,
                    length :15,
                    lineStyle: {
                        color: 'auto'
                    }
                },
                pointer: {
                    width: 4
                },
                title: {
                    show: true,
                    offsetCenter: [0, '30%'],
                    textStyle: {
                        fontWeight: 'bolder'
                    }
                },
                radius: [0, '100%'],
                detail: {
                    formatter: '{value}',
                    textStyle: {
                        color: 'auto',
                        fontWeight: 'bolder'
                    }
                },
                data: [{value: 0, name: '主轴倍率'}]
            }
        ]
    };
    mdcEquiMoniGaugeOption2 = $.extend(true, {}, mdcEquiMoniGaugeOption1);
    mdcEquiMoniGaugeOption2.series[0].name = '进给倍率';
    mdcEquiMoniGaugeOption2.series[0].data[0].name = '进给倍率';
    mdcEquiMoniGaugeOption3 = $.extend(true, {}, mdcEquiMoniGaugeOption1);
    mdcEquiMoniGaugeOption3.series[0].name = '主轴负载';
    mdcEquiMoniGaugeOption3.series[0].data[0].name = '主轴负载';

    getEquipmentState();
    equipmentInterval = setInterval(getEquipmentState,5 * 1000);
});

function getEquipmentState() {
    if (!$("#equipment_monitor_div").length) {
        clearInterval(equipmentInterval);
        return;
    }

//	$.ajax({
//        url : "/equipmentMonitor/equipmentallstate",
//        dataType : "json",
//        success : function(data, textStatus, xhr) {
			
//            $.each(data,function(index, value) {
//                equipmentAll.find("#equipment_" + value.EquipmentID).attr("src", setEquipmentIcon(value.Oporation))
//                    .closest("a").attr({"dataopo":value.Oporation,"dataname":value.EquipmentName});
//            });
//        }
//    })
    $.ajax(ajaxSettings({
        url : "/equipmentMonitor/equipmentallstate",
        dataType : "json",
        success : function(data) {
			if (data.success != undefined) {
				clearInterval(equipmentInterval);
				return
			}
            $.each(data,function(index, value) {
                equipmentAll.find("#equipment_" + value.EquipmentID).attr("src", setEquipmentIcon(value.Oporation))
                    .closest("a").attr({"dataopo":value.Oporation,"dataname":value.EquipmentName});
            });
        }
    }))
}

$("#equipment_monitor_div a").bind('click',function() {
    var $this = $(this);

//    console.log($this.attr("dataopo"));
    if (parseInt($this.attr("dataopo"))) {
        equipmentDetail($this.attr("datatype"), $this.attr("dataname"));
        var init = setInterval(function () {
            equipmentDetail($this.attr("datatype"), $this.attr("dataname"));
        }, 3000);
        $('#mdcEquiInfo').show().dialog({
            title: '设备信息',
            width: 900,
            height: 600,
            modal: true,
            onClose: function () {
                clearInterval(init);
            }
        });
    } else {
        $.messager.alert('提示信息', '设备处于关机状态！');
    }
})

/*设备状态颜色图标*/
function setEquipmentIcon(oporation){
    switch(oporation){
        case 1 :
        case 2 :
            return '/static/images/mdc/yellow.png';
        case 3 :
            return '/static/images/mdc/green.png';
        case 22 :
            return '/static/images/mdc/red.png';
        default :
            return '/static/images/mdc/off.png';
    }
}

function equipmentDetail(equipmentID, equipmentName) {
    $('#inpMdcLogEquiId').textbox('setValue', equipmentID);
    $('#inpMdcLogEquiName').textbox('setValue', equipmentName);
    $('#inpMdcStatChartEquiId').textbox('setValue', equipmentID);
    $('#inpMdcStatChartEquiName').textbox('setValue', equipmentName);
    $("#mdcEquiMoni-equipmentID").textbox('setValue', equipmentID);//设备编号
    $("#mdcEquiMoni-equipmentName").textbox('setValue', equipmentName);//设备名称
    $.post('/equipmentMonitor/findByEquipmentId', {equipmentId: equipmentID}, function (data) {
        if (data[0] != null) {
            data = data[0];
            // 主轴倍率
            mdcEquiMoniGaugeOption1.series[0].data[0].value = data.spindlebeilv || 0;
            // 进给倍率
            mdcEquiMoniGaugeOption2.series[0].data[0].value = data.feedbeilv || 0;
            // 主轴负载
            mdcEquiMoniGaugeOption3.series[0].data[0].value = data.spindleload || 0;
            mdcEquiMoniGauge1.setOption(mdcEquiMoniGaugeOption1);
            mdcEquiMoniGauge2.setOption(mdcEquiMoniGaugeOption2);
            mdcEquiMoniGauge3.setOption(mdcEquiMoniGaugeOption3);
            //$("#mdcEquiMoni-equipmentID").textbox('setValue', equipmentID);//设备编号
            //$("#mdcEquiMoni-equipmentName").textbox('setValue', equipmentName);//设备名称
            //$("#mdcEquiMoni-continuousSys").textbox('setValue', value.equipmentID);//控制系统
            $("#mdcEquiMoni-executingcode").textbox('setValue', data.Programnumber || "");//当前执行程序
            $("#mdcEquiMoni-spindlebeilv").textbox('setValue', data.spindlebeilv);//主轴倍率
            $("#mdcEquiMoni-feedbeilv").textbox('setValue', data.feedbeilv);//进给倍率
            $("#mdcEquiMoni-spindleload").textbox('setValue', data.spindleload);//主轴负载
            $("#mdcEquiMoni-spindlespeed").textbox('setValue', lespeed(data.spindlespeed));//主轴转速 mdcEquiMoni-spindlespeed
            $("#mdcEquiMoni-feedrate").textbox('setValue', lespeed(data.feedrate));//进给速度
            //$("#mdcEquiMoni-ALRMstate").textbox('setValue', data.data.ALRMstate);//报警
        }
    });
}

/*
 主轴转速，进给速度单位
 0 : mm/min
 1 : inch/min
 2 : rpm
 3 : mm/rev
 4 : inch/rev
 */
function lespeed(speed) {
    if(speed == null){
        return '';
    }
    var tag = speed.substring((speed.indexOf('(') + 1), (speed.length - 1));
    var s = '';
    switch (tag) {
        case '0':
            s = 'mm/min';
            break;
        case '1' :
            s = 'inch/min';
            break;
        case  '2' :
            s = 'rpm';
            break;
        case  '3' :
            s = 'mm/rev';
            break;
        case  '4' :
            s = 'inch/rev';
            break;
    }
    return speed.substring(0, (speed.length - 3)) + s;
}
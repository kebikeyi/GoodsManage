var chm = 0;
var menu =
[
	['我的主页','index.html'],
	['快递查询','express.html'],
	['我的资料','myinfo.html'],
	[
		['基本信息','user.html'],
		['修改密码','pwd.html'],
	],
	['系统帮助','help.html']
];

function $(id) {
	return document.getElementById(id);
}

var currentfile = location.href.substr(location.href.lastIndexOf('/') + 1);
function docmenu(mdiv) {
	var showtype=1;
	var returnstr = '';
	if(showtype && chm) {
		document.body.style.background = 'none';
		$('wrap').style.paddingLeft = 0;
		return;
	}
	var menucount = 0;
	var tabon;
	for(var i in menu) {
		if(typeof(menu[i][0]) == 'object') {
			if(showtype) {
				returnstr += '<div class="subinfo" id="menu' + menucount + '" style="display: ">';
				for(var k in menu[i]) {
					tabon = '';
					if(currentfile == menu[i][k][1]) {
						tabon = 'tabon ';
					}
					if(!menu[i][k][1]) {
						menu[i][k][1] = '';
					}
					//returnstr += '<a target="main" onclick="'+"$('admincpnav').innerHTML='当前功能:<span >->>"+menu[i][k][0]+"&nbsp;<a href=http://baidu.com>刷新显示</a></span>';"+'" class="' + tabon + 'sidelist" href="' + menu[i][k][1] + '">' + menu[i][k][0] + '</a>';
					returnstr += '<a target="main" onclick="navshow(\''+menu[i][k][0]+'\',\''+menu[i][k][1]+'\')" class="' + tabon + 'sidelist" href="' + menu[i][k][1] + '">' + menu[i][k][0] + '</a>';
				
				}
				returnstr += '</div>';
			}
		} else {
			tabon = '';
			if(!menu[i][1]) {
				menu[i][1] = '';
			}
			if(showtype) {
				menucount++;
				if(currentfile == menu[i][1]) {
					tabon = 'tabon ';
				}
				returnstr += '<a target="main"   onclick="navshow(\''+menu[i][0]+'\',\''+menu[i][1]+'\')" class="' + tabon + 'sideul"';
				if(menu[i][1] != '') {
					returnstr += ' href="' + menu[i][1] +'"';
				}
				returnstr += '><em class="shrink" onclick="collapse(this, \'menu' + menucount + '\');return false">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</em>' + menu[i][0] + '</a>';
			} else {
				returnstr += '<li><a target="main" ';
				if(menu[i][1] != '') {
					returnstr += ' href="' + menu[i][1] +'"';
				}
				returnstr += '>' + menu[i][0] + '</a></li>';
			}
		}
	}
	if(showtype) {
		//document.write('<div class="side" style="height: 400px;">' + returnstr + '</div>');
		$(mdiv).innerHTML = '<div class="side" style="height: 400px;padding-left:10px;">' + returnstr + '</div>';
	} else {
		return '<ul>' + returnstr + '</ul>';
	}
}
/*
function showmenu(ctrl) {
	ctrl.className = ctrl.className == 'otherson' ? 'othersoff' : 'otherson';
	var menu = parent.document.getElementById('toggle');
	if(!menu) {
		menu = parent.document.createElement('div');
		menu.id = 'toggle';
		menu.innerHTML = docmenu(0);
		var obj = ctrl;
		var x = ctrl.offsetLeft;
		var y = ctrl.offsetTop;
		while((obj = obj.offsetParent) != null) {
			x += obj.offsetLeft;
			y += obj.offsetTop;
		}
		menu.style.left = x + 'px';
		menu.style.top = y + ctrl.offsetHeight + 'px';
		menu.className = 'togglemenu';
		menu.style.display = '';
		parent.document.body.appendChild(menu);
	} else {
		menu.style.display = menu.style.display == 'none' ? '' : 'none';
	}
}
*/
function collapse(ctrlobj, showobj) {
	if(!$(showobj)) {
		return;
	}
	if($(showobj).style.display == '') {
		ctrlobj.className = 'spread';
		$(showobj).style.display = 'none';
	} else {
		ctrlobj.className = 'shrink';
		$(showobj).style.display = '';
	}
}

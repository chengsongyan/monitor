{{define "navbar"}}
<a class="navbar-brand" href="/">监控系统</a>
<div>
	<ul class="nav navbar-nav">
		<li {{if .IsHome}}class="active"{{end}}><a href="/">主机状态</a></li>
		<li {{if .IsAddHost}}class="active"{{end}}><a href="/addhosts">添加主机</a></li>
<!--		<li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>-->
	</ul>
</div>

<div class="pull-right">
	<ul class="nav navbar-nav">
	{{if .IsLogin}}
	<li><a href="/login?exit=true">退出登录</a></li>
	{{else}}
	<li><a href="/login">管理员登录</a></li>
	{{end}}
	</ul>
</div>

{{end}}
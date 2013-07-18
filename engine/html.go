package engine

// THIS FILE IS AUTO GENERATED FROM gui.html
// EDITING IS FUTILE

const templText = `
<!DOCTYPE html>
<html>

<head>

	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">

	<title>mx3 GUI</title>

	<style media="all" type="text/css">

		body  { margin: 20px; font-family: Ubuntu, Arial, sans-serif; font-size: 14px; color: #444444}
		h1    { font-size: 22px; font-weight: normal; color: black}
		h2    { font-size: 18px; font-weight: normal; color: black}
		img   { margin: 15px; }
		table { border:10px; border-collapse: collapse; margin:10px}
		tr:nth-child(even) { background-color: white; }
		tr:nth-child(odd)  { background-color: #F0F0FF; }
		td    { padding: 1px 5px;}
		hr    { border-style: none; border-top: 1px solid #CCCCCC; }
		a     { color: #375EAB; text-decoration: none; }

		div       { margin-left: 20px; margin-top: 5px; margin-bottom: 20px; }
		div#footer{ color:gray; font-size:14px; border:none}

		.ErrorBox { color: red; font-weight: bold; } 
		.TextBox  { border:solid; border-color:#BBBBBB; border-width:1px; padding-left:4px;}
	</style>

	<script>
		function toggle(id) {
	       var el = document.getElementById(id);
	       if(el.style.display != 'none'){
	          el.style.display = 'none';
			} else {
	          el.style.display = 'block';
			}
	    }
	</script>

	{{.JS}}

</head>


<body>

	<h1> mx 3 GUI </h1> 
		<p> {{.ErrorBox}} </p>
		<p> {{.Span "log"}} </p>
	<hr/>

	
	<h2> geometry </h2><div>

		<table>
			<tr> <td><b>gridsize: </b></td> <td>{{.Span "nx"}} </td> <td> &times; {{.Span "ny"}}</td> <td> &times; {{.Span "nz"}}                </td> </tr>
			<tr> <td><b>cellsize: </b></td> <td>{{.Span "cx"}} </td> <td> &times; {{.Span "cy"}}</td> <td> &times; {{.Span "cz"}} nm<sup>3</sup> </td> </tr>
			<tr> <td><b>worldsize:</b></td> <td>{{.Span "wx"}} </td> <td> &times; {{.Span "wy"}}</td> <td> &times; {{.Span "wz"}} nm<sup>3</sup> </td> </tr>
		</table>
	</div><hr/>


	<h2> solver </h2><div>
		<table>
			<tr> <td> <b>status:</b> {{.Span "solverstatus" "initializing"}} </td><td>{{.Button "break"}}        </td></tr>
			<tr> <td> {{.Button "run"}}               </td><td>{{.NumBox "runtime" 1e-9}} </td></tr> 
			<tr> <td> {{.Button "steps"}}             </td><td>{{.NumBox "runsteps" 1000}}</td></tr>
			<tr> <td> {{.Button "relax"}}             </td><td>                           </td></tr>
		</table>

		<table>
			<tr> <td><b>step:    </b></td><td>{{.Span "step"}}          </td><td><b>time:  </b></td> <td>{{.Span "time"}} s       </td></tr>
			<tr> <td><b>dt:      </b></td><td>{{.NumBox "dt" 1e-15}}   s</td><td>   {{.CheckBox "fixdt" "fixdt" false}}           </td></tr>
			<tr> <td><b>mindt:   </b></td><td>{{.NumBox "mindt" 0}}s    </td><td><b>maxdt: </b></td><td>{{.NumBox "maxdt" 0}}s    </td></tr>
			<tr> <td><b>err/step:</b></td><td>{{.Span "err"}}           </td><td><b>maxerr:</b></td><td>{{.NumBox "maxerr" 1e-5}} </td></tr>
		</table>
	</div><hr/>


	<h2> process </h2><div>
		<table>
			<tr> <td><b>host:     </b></td> <td>{{.Span "hostname"}} </td> <td> </td>                           </tr>
			<tr> <td><b>pid:      </b></td> <td>{{.Span "pid"}}      </td> <td> {{.Button "kill" "Kill"}} </td> </tr>
			<tr> <td><b>gpu:      </b></td> <td>{{.Span "gpu"}}      </td> <td>  </td>                          </tr>
			<tr> <td><b>walltime: </b></td> <td>{{.Span "walltime"}} </td> <td>  </td>                          </tr>
		</table>
	</div><hr/>

	<div id="footer">
		<br/>
		<center>
			mumax 3<br/>
			Copyright 2012-2013 <a href="mailto:a.vansteenkiste@gmail.com">Arne Vansteenkiste</a>,
			<a href="http://dynamat.ugent.be">Dynamat group</a>, Ghent University, Belgium.<br/>
			You are free to modify and distribute this software under the terms of the
			<a href="http://www.gnu.org/licenses/gpl-3.0.html">GPLv3 license</a>.
		</center>
	</div>

</body>
</html>
`
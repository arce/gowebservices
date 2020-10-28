<?php 
header('Access-Control-Allow-Origin: *');
header("Access-Control-Allow-Headers: X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Request-Method");
header("Access-Control-Allow-Methods: GET, POST, OPTIONS, PUT, DELETE");
header("Allow: GET, POST, OPTIONS, PUT, DELETE");
?>
<html>
<head>
  <script type="text/javascript" src="js/fetchival.min.js"></script>
  <script type="text/javascript" src="app.js"></script>
</head>
<body>
  <script>
	  test();
	  alert("This is alert box!");
  </script>
</body>
</html>
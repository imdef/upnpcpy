$pipe = new-object System.IO.Pipes.NamedPipeClientStream("\\.\pipe\Wulf");


$pipe.Connect(); 


 
$sw = new-object System.IO.StreamWriter($pipe);
$sw.WriteLine("Go"); 
$sw.WriteLine("start abc 123"); 
$sw.WriteLine('exit'); 
 
$sw.Dispose(); 
$pipe.Dispose();
$pipe=new-object System.IO.Pipes.NamedPipeServerStream("\\.\pipe\Wulf");
'Created server side of "\\.\pipe\Wulf"'
$pipe.WaitForConnection(); 
"connection." 

$sr = new-object System.IO.StreamReader($pipe); 
while (($cmd= $sr.ReadLine()) -ne 'exit') 
{
 $cmd
}; 
 
$sr.Dispose();
$pipe.Dispose();
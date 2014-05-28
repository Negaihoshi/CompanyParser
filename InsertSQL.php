<?php
    $link = mysql_connect('localhost', 'root', 'mysql');
    if (!$link) {
        die('Could not connect: ' . mysql_error());
    }
    mysql_query("SET NAMES 'utf8'");
    mysql_select_db("company", $link);
    echo 'Connected successfully<br>';

    //mysql_close($link);

    $row = 1;
    if (($handle = fopen("index.csv", "r")) !== FALSE) {
        while (($data = fgetcsv($handle, 1000, ",")) !== FALSE) {
            if($row!=1){
                //echo $data[0] . " " . $data[1] . " " . $data[2] . "<br />\n";
                $data[0] = (string)$data[0];
                $data[0] = htmlspecialchars(stripslashes(trim($data[0])), ENT_QUOTES);
                $data[1] = htmlspecialchars(stripslashes(trim($data[1])), ENT_QUOTES);
                $data[2] = htmlspecialchars(stripslashes(trim($data[2])), ENT_QUOTES);
                if($data[2]=="") echo "<hr>".$data[0].$data[1].$data[2]."<hr>";
                $sql = "INSERT INTO companyIndex(Id,Type,Name) VALUES ('$data[0]','$data[1]','$data[2]')";
                mysql_query($sql) or die('<br>Insert data fail: '.mysql_error());

            }else{
                $row+=1;
            }
        }
        fclose($handle);
    }
?>

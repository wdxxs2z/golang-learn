<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>这是我的上传页面</title>
</head>
<body>
    <form enctype="multipart/form-data" action="/upload" method="post">
        上传文件:<input type="file" name="uploadFile" />
        <input type="submit" name="上传">
    </form>
</body>
</html>
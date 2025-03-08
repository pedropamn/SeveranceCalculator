<?php

if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $nome = htmlspecialchars($_POST['nome']);
    $email = filter_var($_POST['email'], FILTER_SANITIZE_EMAIL);
    $mensagem = htmlspecialchars($_POST['mensagem']);

    // Destinatário do e-mail
    $para = "contato@dominio.com"; 

    // Assunto do e-mail
    $assunto = "Novo contato de $nome";

    // Corpo do e-mail
    $corpo = "Nome: $nome\n";
    $corpo .= "E-mail: $email\n";
    $corpo .= "Mensagem:\n$mensagem\n";

    // Cabeçalhos do e-mail
    $headers = "From: $email\r\n";
    $headers .= "Reply-To: $email\r\n";
    $headers .= "Content-Type: text/plain; charset=UTF-8\r\n";

    // Enviar e-mail
    if (mail($para, $assunto, $corpo, $headers)) {
        echo "<h1 style='color: green;'>Mensagem enviada com sucesso!</h1>\n\n<a style='font-size:30px;text-align:center; margin:auto;' href='javascript:history.back();'>Voltar</a>";
    } else {
        echo "<h1 style='color: red;'>Erro ao enviar a mensagem. Tente novamente</h1>\n\n<a style='font-size:30px;text-align:center; margin:auto;' href='javascript:history.back();'>Voltar</a>";
    }
} else {
    echo "Método de requisição inválido.";
}
?>

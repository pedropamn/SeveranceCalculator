<!DOCTYPE html>
<html lang="en" >
<head>
  <meta charset="UTF-8">
  <title>Featured Tabs</title>
  <link rel='stylesheet' href='../assets/css/ribbon.css'> <!-- https://codepen.io/magnusriga/pen/aKopeG -->
  <link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css'>
<link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css'>
<link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,200,300,700'><link rel="stylesheet" href="./style.css">

</head>
<body>
<!-- partial:index.partial.html -->
<div class="container"> 
<section id="fancyTabWidget" class="tabs t-tabs">
        <ul class="nav nav-tabs fancyTabs" role="tablist">
        
                    <li class="tab fancyTab active" style="width:100%;">
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>	
                        <a id="tab0" href="#tabBody0" role="tab" aria-controls="tabBody0" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-fire"></span><span class="title-tab">Rescisão</span></a>
                    	<div class="whiteBlock"></div>
                    </li>
                    <!--
                    <li class="tab fancyTab">
					  <div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab1" href="#tabBody1" role="tab" aria-controls="tabBody1" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-firefox"></span><span class="title-tab">Férias</span></a>
                        <div class="whiteBlock"></div>
                    </li>
                    
                    <li class="tab fancyTab">
					<div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab2" href="#tabBody2" role="tab" aria-controls="tabBody2" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-envira"></span><span class="title-tab">Décimo Terc.</span></a>
                        <div class="whiteBlock"></div>
                    </li>
                    
                    <li class="tab fancyTab">
					<div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab3" href="#tabBody3" role="tab" aria-controls="tabBody3" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-mortar-board"></span><span class="title-tab">FGTS</span></a>
                        <div class="whiteBlock"></div>
                    </li> 
                         
                    <li class="tab fancyTab">
					<div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab4" href="#tabBody4" role="tab" aria-controls="tabBody4" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-stack-overflow"></span><span class="title-tab">Seguro Desempr.</span></a>
                        <div class="whiteBlock"></div>
                    </li>
                    
                    <li class="tab fancyTab">
					<div class="ribbon blue"><span>Em breve</span></div>
                    <div class="arrow-down"><div class="arrow-down-inner"></div></div>
                        <a id="tab5" href="#tabBody5" role="tab" aria-controls="tabBody5" aria-selected="true" data-toggle="tab" tabindex="0"><span class="fa fa-question-circle"></span><span class="title-tab">Adic. Noturno</span></a>
                        <div class="whiteBlock"></div>
                    </li>
					-->
        </ul>
        <div id="myTabContent" class="tab-content fancyTabContent" aria-live="polite">
                    <div class="tab-pane  fade active in" id="tabBody0" role="tabpanel" aria-labelledby="tab0" aria-hidden="false" tabindex="0">
                        <div>
                        	<div class="row">
                            	
                                <div class="col-lg-12">
									<form action="#" method="post" class="php-email-form aos-init aos-animate" data-aos="fade-up" data-aos-delay="200">
										<div class="row gy-4">
										
											<div class="col-md-12">
												<h2>Rescisão</h2>
											   
											</div>
											
											<!-- Input para Data de Admissão -->
											<div class="col-md-6">
												<label for="diasTrabalhados" class="form-label">Data de Admissão:</label>
												<input type="date" class="form-control" id="dataAdmissao" name="dataAdmissao" placeholder="" required>
											</div>
											
											<!-- Input para Data de Demissão -->
											<div class="col-md-6">
												<label for="diasTrabalhados" class="form-label">Data de Demissão:</label>
												<input type="date" class="form-control" id="dataDemissao" name="dataDemissao" placeholder="" required>
											</div>

											<!-- Input para Salário Mensal -->
											<div class="col-md-6">
												<label for="salarioMensal" class="form-label">Salário Mensal</label>
												<input type="number" class="form-control" id="salarioMensal" name="salarioMensal" placeholder="Ex.: 3000.00" step="0.01" required>
											</div>											

											<!-- Input para Tipos de Rescisão -->
											<div class="col-md-6">
												<label for="tipoRescisao" class="form-label">Tipo de Rescisão:</label>
												<select id="tipoRescisao" class="form-control">
													<option value="justa_causa">Justa causa</option>
													<option value="sem_justa_causa">Sem justa causa</option>
													<option value="se_demitiu">Se demitiu</option>
												</select>
											</div>

											<!-- Input para meses acumulados de férias-->
											<div class="col-md-6">
												<label for="mesesAcumuladosFerias" class="form-label">Meses acumulados de férias</label>
												<input type="number" class="form-control" id="mesesAcumuladosFerias" name="mesesAcumuladosFerias" placeholder="Ex.: 8" required>
											</div>
											

											<div class="col-md-12 text-center">
												<div class="loading" style="display:none;">Carregando</div>
												<div class="error-message"></div>
												<div class="sent-message" style="display:none;">Sua mensagem foi enviada. Obrigado!</div>
												
												<br>
												
												<div class="bloco_resultado" style="display:none;">
													<span style="font-size: 20px;">Você tem direito a receber: </span>
													<div class="resultado" style="font-size: 40px; color: green;"></div>
												</div>
												<button class="btn btn-success btn-lg btnCalcular">Calcular</button>
											</div>

										</div>
									</form>
								</div>
                                
                            </div>
                        </div>
                    </div>
					<!--
                    <div class="tab-pane  fade" id="tabBody1" role="tabpanel" aria-labelledby="tab1" aria-hidden="true" tabindex="0">
                        <div class="row">
                            	
                                <div class="col-md-12">
                        			<h2>Em breve</h2>
                                    <p>Em breve</p>
                                   
                                </div>
                            </div>
                    </div>
					
                    <div class="tab-pane  fade" id="tabBody2" role="tabpanel" aria-labelledby="tab2" aria-hidden="true" tabindex="0">
                        <div class="row">
							<div class="col-md-12">
								<h2>Em breve</h2>
								<p>Em breve</p>                                  
							</div>
						</div>
                    </div>
                    <div class="tab-pane  fade" id="tabBody3" role="tabpanel" aria-labelledby="tab3" aria-hidden="true" tabindex="0">
						 <div class="row">
							<div class="col-md-12">
								<h2>Em breve</h2>
								<p>Em breve</p>                                  
							</div>
						</div>
                    </div>
                    <div class="tab-pane  fade" id="tabBody4" role="tabpanel" aria-labelledby="tab4" aria-hidden="true" tabindex="0">
                     <div class="row">
							<div class="col-md-12">
								<h2>Em breve</h2>
								<p>Em breve</p>                                  
							</div>
						</div>
                    </div>
                    <div class="tab-pane  fade" id="tabBody5" role="tabpanel" aria-labelledby="tab5" aria-hidden="true" tabindex="0">
					
                     <div class="row">
							<div class="col-md-12">
								<h2>Em breve</h2>
								<p>Em breve</p>                                  
							</div>
						</div>
                    </div>
					-->
        </div>

    </section>
</div>
<!-- partial -->
<script src='//cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js'></script>
<script src='https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js'></script><script  src="./script.js"></script>
<script src='../assets/js/main.js'></script>

<script>
$(document).ready(function() {
    $('.btnCalcular').click(function(e) {
        e.preventDefault(); // Impede o recarregamento da página

        // Validação dos campos
        if ($('#dataAdmissao').val() === "" || 
            $('#dataDemissao').val() === "" || 
            $('#salarioMensal').val() === "" || 
            $('#tipoRescisao').val() === "" ||
            $('#mesesAcumuladosFerias').val() === "") {
            alert("Por favor, preencha todos os campos.");
            return; // Interrompe a função caso algum campo esteja vazio
        }

        const dados = {
            dataAdmissao: $('#dataAdmissao').val(),
            dataDemissao: $('#dataDemissao').val(),
            salarioMensal: $('#salarioMensal').val(),
            tipoRescisao: $('#tipoRescisao').val(),
            mesesAcumuladosFerias: $('#mesesAcumuladosFerias').val()
        };

        $.ajax({
            url: '/', // Endereço da requisição POST
            method: 'POST',
            data: dados, // Envia os dados como form data
            success: function(response) {
                console.log('sucesso');
                $('.bloco_resultado').css('display', 'block');
                $('.resultado').text('R$ ' + response.valor);
            },
            error: function(err) {
                console.log('erro');
                alert("Erro ao calcular. Tente novamente");
            }
        });
    });
});



</script>

</body>
</html>

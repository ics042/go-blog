$(function(){
	window.sr = ScrollReveal();
	sr.reveal('.service-list', { reset: true, duration: 1000, delay: 200 }, 150);
	
	$('#top-arrow-down').on('click', function ( e ) {
		
		target = $("#service").offset();
		target = target.top;
		if ( target ) {
			$('html, body').stop().animate({scrollTop: target}, 700, "linear");
		}
	});
	
	$('#main-nav li a').bind('click',function(event){
			var $anchor = $(this);
			console.log($anchor);
			console.log($($anchor.attr('href')).offset().top);
			$('html, body').stop().animate({
				scrollTop: $($anchor.attr('href')).offset().top
			}, 1500,'easeInOutExpo');
			event.preventDefault();
		});
		
	var $container = $('.portfolioContainer'),
	  $body = $('body'),
	  colW = 375,
	  columns = null;

	$container.isotope({
	// disable window resizing
	resizable: true,
	masonry: {
	  columnWidth: colW
	}
	});
  
	$('.portfolioFilter a').click(function(){
		$('.portfolioFilter .current').removeClass('current');
		$(this).addClass('current');

		var selector = $(this).attr('data-filter');
		$container.isotope({
			
			filter: selector,
		 });
		 return false;
	});
	
});
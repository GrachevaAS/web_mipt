function add_new_task(text) {
	$("ul").append("<li>");
	last = $("ul").children().last();
	last.append($("<span>").text(text));
	last.append($("<button class='delete'>").text("Удалить"))
	$(".delete").click(function () {
		$(this).parent().remove();
	})
}

$(document).ready(function() {
	$('#root').append("<ul>");
	add_new_task("Сделать задание #3 по web-программированию");
	$('#root').append("<input id='add_task_input'>");
    $('#root').append($("<button id = 'add_task'>").text('Добавить задачу'))    
	$("#add_task").click(function() {
		add_new_task($("#add_task_input").val());
		document.getElementById("add_task_input").value = "";
	})
})
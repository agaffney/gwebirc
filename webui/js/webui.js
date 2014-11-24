function get_connections() {
	$.getJSON('/api/connections', function(data) {
		alert(JSON.stringify(data))
//		alert("You are connected to host " + data[0]['host'] + " with nick " + data[0]['nick'])
	});
}

function test_breadcrumbs() {
	var random_items = ["apple","banana","car","donut"].sort(function() {
		return (Math.round(Math.random())-0.5);
	});
	var data = [
		{ "text": "Home", "url": "#" },
	];
	for (var x = 0; x < random_items.length; x++) {
		var item = random_items[x];
		data.push({ "text": item, "url": "#" });
	}
	update_breadcrumbs(data);
}

function test_events() {
	$.getJSON('/api/events/80', function(data) {
		console.log(JSON.stringify(data))
	});
}

function update_breadcrumbs(items) {
	var br = $('.breadcrumb')
	br.empty()
	var items_len = items.length;
	for (var i = 0; i < items_len - 1; i++) {
		var item = items[i];
		br.append('<li><a href="' + item['url'] + '">' + item['text'] + '</a></li>')
	}
	// Handle last item differently, since it's the current page
	br.append('<li class="active">' + items[items_len - 1]['text'])
}

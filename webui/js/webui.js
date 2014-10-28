function get_connections() {
	$.getJSON('/api/connections', function(data) {
		alert("You are connected to host " + data[0]['host'] + " with nick " + data[0]['nick'])
	});
}

// display back button
var pathname = window.location.pathname
if (pathname != "/") {
	document.getElementById("back").style.display = "inline"
} else {
	document.getElementById("back").style.display = "none"
}

function setHref(dir) {
	if (window.location.href.charAt(window.location.href.length-1) != "/") {
		window.location.href += "/" + dir 
	} else {
		window.location.href += dir + "/"
	}
};

function serveFile(file_path) {
	var pathname = window.location.pathname
	if (pathname.charAt(pathname.length-1) != "/") {
		pathname += "/"
	}
	window.location.href = "/api/v1/file/serve/" + pathname + file_path 
};

function back() {
	var pathname = window.location.pathname
	if (pathname == "/") {
		return
	}
	var pathname_split = pathname.split("/")

	pathname_split.pop() // remove one element
	if (pathname.charAt(pathname.length-1) == "/") { // remove two elements if pathname ends with '/'
		pathname_split.pop()
	}
	window.location.href = pathname_split.join("/") + "/"
}

function openFileModal(file_path) {
	var pathname = window.location.pathname
	if (pathname.charAt(pathname.length-1) != "/") {
		pathname += "/"
	}
	
	var file_location = "/api/v1/file/serve/" + pathname + file_path
	console.log(file_location)

	var options = {
		hashTracking: false
	}
	var inst = $('[data-remodal-id=file-modal]').remodal(options)

	// set values
	$("#file-modal-title").text(file_path)
	inst.open()
}


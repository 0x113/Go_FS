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

function newFolderModal() {
	var options = {
		hashTracking: false
	}
	var inst = $("[data-remodal-id=new-folder-modal]").remodal(options)
	inst.open()
	$("#new-folder-form").submit((e) => {
		createNewDir(e, inst)
	})
}

createNewDir = (e, inst) => {
	let pathname = window.location.pathname
	let form_data = {
		"path": decodeURI(pathname),
		"dir_name": $("input[name=folder-name]").val()
	}
	$.ajax({
		type: "POST",
		url: "/api/v1/directory/new",
		data: form_data,
		dataType: "json",
		encode: true,
		success: (response) => {
			console.log("res", response)
			inst.close()
			location.reload()
		},
		error: (jqXHR, textStatus, errorThrown) => {
			var err = jqXHR.responseJSON["error"]
			new Noty({
				type: "error",
				theme: "nest",
				layout: "topRight",
				text: "Unable to create new directory",
				timeout: 5000,
				progressBar: true
			}).show();
		},
	})
	e.preventDefault()
}

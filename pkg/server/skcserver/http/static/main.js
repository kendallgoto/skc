async function update() {
	const body = document.codeEditor.getValue();
	res = await fetch("/", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({ code: body })
	})
	let resJson = await res.json()
	document.codeEditor.operation(function () {
		for (error of document.codeErrors) {
			error.clear()
		}
		document.codeErrors = []
		if (resJson.errors) {
			// errors exist
			document.transformViewer.display.wrapper.classList.add("invalid")
			document.executionResult.classList.add("invalid")
			document.transformViewer.setValue("")
			document.executionResult.innerHTML = ""
			// add lint annotations
			for (error of resJson.errors) {
				let errorEl = document.createElement("div");
				errorEl.appendChild(document.createTextNode(error.msg));
				errorEl.className = "lint-error";
				let widget = document.codeEditor.addLineWidget(error.line - 1, errorEl)
				console.log(widget)
				if (widget)
					document.codeErrors.push(widget)
			}
		} else {
			document.transformViewer.display.wrapper.classList.remove("invalid")
			document.executionResult.classList.remove("invalid")
			document.transformViewer.setValue(resJson.output)
			document.executionResult.innerHTML = resJson.execution
		}
	})
	console.log(resJson)
}
async function initWithCode(code) {
	const codeEditor = CodeMirror(document.getElementById("codemirror-code"),
		{
			tabSize: 4,
			indentWithTabs: true,
			lineNumbers: true,
			lineWrapping: true,
			value: code
		}
	)
	const transformViewer = CodeMirror(document.getElementById("codemirror-transform"),
		{
			mode: "text/x-python",
			tabSize: 4,
			indentWithTabs: true,
			lineNumbers: true,
			lineWrapping: false,
			readOnly: true,
			value: ""
		}
	)
	document.codeEditor = codeEditor
	document.transformViewer = transformViewer
	document.transformViewer.display.wrapper.classList.add("readonly")
	document.executionResult = document.getElementById("execute-result")
	document.codeErrors = []
	document.codeEditor.on("change", function (cm, change) {
		// debounce
		if (cm.ChangeBuffer) {
			clearTimeout(cm.ChangeBuffer)
			cm.ChangeBuffer = false;
		}
		cm.ChangeBuffer = setTimeout(async function () {
			cm.ChangeBuffer = false;
			console.log("debounce fire")
			await update()
		}, 200)
	})
	await update()
}

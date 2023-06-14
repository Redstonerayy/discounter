const codedisplay = document.querySelector(".code-display");
const copy = document.querySelector(".new-code > img");
const bnew = document.querySelector(".b-new");

const codeinput = document.querySelector(".input-field");
const result = document.querySelector(".check-inval-reset-result");
const bcheck = document.querySelector(".b-check");
const binval = document.querySelector(".b-inval");
const breset = document.querySelector(".b-reset");

copy.addEventListener("click", () => {
	navigator.clipboard.writeText(codedisplay.textContent);
});

bnew.addEventListener("click", async () => {
	const response = await fetch("/new-code");
	const json = await response.json();
	codedisplay.textContent = json.newcode.toString().padStart(16, "0");
});

bcheck.addEventListener("click", async () => {
	const response = await fetch("/check-code", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ code: Number(codeinput.value) }),
	});

	const json = await response.json();

	if(json.codestatus == "not found"){
		result.textContent = "Code not found!"
	} else if (json.codestatus == "valid"){
		result.textContent = "Code is valid!"
	} else if (json.codestatus == "invalid"){
		result.textContent = "Code is invalid!"
	}
});

binval.addEventListener("click", async () => {
	const response = await fetch("/inval-code", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ code: Number(codeinput.value) }),
	});

	const json = await response.json();
	
	if(json.codestatus == "not found"){
		result.textContent = "Code not found!"
	} else if (json.codestatus == "invalid"){
		result.textContent = "Code is now invalid!"
	}
});

breset.addEventListener("click", async () => {
	const response = await fetch("/reset-code", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ code: Number(codeinput.value) }),
	});

	const json = await response.json();
	
	if(json.codestatus == "not found"){
		result.textContent = "Code not found!"
	} else if (json.codestatus == "valid"){
		result.textContent = "Code is now valid again!"
	}
});

async function shorten() {
    try {
        new URL(url.value);
    } catch (e) {
        return alert(e.toString().split("constructor:")[1])
    }

    let params = new URLSearchParams()
    params.set("url", url.value)
    if (slug.value !== "") {
        params.set("slug", slug.value)
    }

    let result = await fetch("/api/shorten?" + params.toString())
    if (result.status !== 200) {
        return alert(await result.text())
    }

    out.value = location.href
    if (!location.href.endsWith("/")) {
        out.value += "/"
    }

    out.value += await result.text()
}

function copyOutput() {
    navigator.clipboard.writeText(out.value)
    alert("copied!")
}
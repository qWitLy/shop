function postToUrl(path, method) {
    method = method || "post"; // Устанавливаем метод отправки.

    var form = document.createElement("form");
    form.setAttribute("method", method);
    form.setAttribute("action", path);

    document.body.appendChild(form);
    form.submit();
}
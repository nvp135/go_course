function setCookie(name, value, options = {}) {
    options = {
        path: '/',
        ...options
    };

    if (options.expires instanceof Date) {
        options.expires = options.expires.toUTCString();
    }

    let updatedCookie = encodeURIComponent(name) + "=" + encodeURIComponent(value);

    for (let optionKey in options) {
    updatedCookie += "; " + optionKey;
    let optionValue = options[optionKey];
    if (optionValue !== true) {
        updatedCookie += "=" + optionValue;
    }
    }

    document.cookie = updatedCookie;
}

function deleteCookie(name) {
    setCookie(name, "", {
        'max-age': -1
    })
}

//document.location = "/testcookies";

window.onload = () => {
    setCookie("appointment", JSON.stringify({
        appointmentDate: "20-01-2021 13:06",
        appointmentStartMn: "1-2",
        appointmentId: 2,
        appointmentUserId: 3
    })
    );
}
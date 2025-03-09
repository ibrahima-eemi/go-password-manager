chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
    if (message.action === "getPassword") {
      fetch("http://localhost:8080/get-password", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ site: message.site, auth: "mon-super-token" })
      })
        .then(response => response.json())
        .then(data => sendResponse({ username: data.username, password: data.password }))
        .catch(error => sendResponse({ error: "Erreur de récupération" }));
    }
    return true;
  });
  
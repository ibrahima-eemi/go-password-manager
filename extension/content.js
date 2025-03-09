chrome.runtime.sendMessage({ action: "getPassword", site: window.location.hostname }, response => {
    if (response && response.username && response.password) {
      let usernameField = document.querySelector("input[type='text'], input[type='email']");
      let passwordField = document.querySelector("input[type='password']");
      if (usernameField && passwordField) {
        usernameField.value = response.username;
        passwordField.value = response.password;
      }
    }
  });
  
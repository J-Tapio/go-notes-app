// Highlights the eye-icon within markdown editor with background color.

window.addEventListener('load', function PlayGroundPageLoaded() {
  if(this.window.location.pathname === "/playground") {
    document.getElementsByClassName("fa-eye")[0].style.backgroundColor = "#d9d9d9";
  }
})


document.addEventListener("DOMContentLoaded", () => {
  const phraseEl = document.querySelector(".content h1");
  if (!phraseEl) return;
  console.log("Rendered excuse:", phraseEl.textContent);
});

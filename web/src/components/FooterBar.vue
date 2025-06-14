<template>
  <footer>
    <div>
      <a
        href="https://github.com/Kylep342/mendel"
        target="_blank"
      >Source</a>
      <br>
    </div>
  </footer>
</template>

<style scoped>
footer {
  /* Core layout with Flexbox */
  display: flex;
  justify-content: space-between; /* This pushes Mendel and Menu to opposite ends */
  align-items: center;

  /* Positioning */
  position: fixed; /* Anchors the header to the top */
  bottom: 0;
  left: 0;
  width: 100%;
  z-index: 1000; /* Ensures it stays on top of other content */

  /* Styling */
  background-color: var(--primary-color);
  color: var(--text-on-primary);
  padding: 0 1.5rem; /* Adds horizontal space on the sides */
  height: 2.5rem; /* Define a fixed height */
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1); /* Optional: adds a subtle shadow */
}
</style>

<script setup lang="ts">
import { useRouter } from "vue-router";
import type { Button } from "@/types/app";
import routes from "@/constants/routes";
import constants from "@/constants/constants";

const router = useRouter();

const menuItems: Array<Button> = [
  { text: constants.BTN_HOME, onClick: () => router.push(routes.ROUTE_INDEX) },
  { text: constants.BTN_ADMIN, onClick: () => router.push(routes.ROUTE_INTERNAL) },
  { text: constants.BTN_PLANT, onClick: () => router.push(routes.ROUTE_PLANT) },
  { text: constants.BTN_PLANT_CULTIVAR, onClick: () => router.push(routes.ROUTE_PLANT_CULTIVAR) },
  { text: constants.BTN_PLANT_SPECIES, onClick: () => router.push(routes.ROUTE_PLANT_SPECIES) },
];

const handleMenuItemClick = (action: () => void, event: Event) => {
  action();

  const detailsElement = (event.target as HTMLElement).closest('details');
  if (detailsElement) {
    detailsElement.removeAttribute('open');
  }
};
</script>

<template>
  <header class="main-header">
    <div class="logo-container">
      <h1>Mendel</h1>
    </div>

    <nav class="main-nav">
      <details class="menu-dropdown">
        <summary>{{ constants.BTN_MENU }}</summary>
        <ul class="menu-items">
          <li v-for="item in menuItems" :key="item.text">
            <a @click.prevent="handleMenuItemClick(item.onClick, $event)">
              {{ item.text }}
            </a>
          </li>
        </ul>
      </details>
    </nav>
  </header>
</template>

<style scoped>
/* Scoped styles ensure these only apply to HeaderBar.vue */
.main-header {
  /* Core layout with Flexbox */
  display: flex;
  justify-content: space-between; /* This pushes Mendel and Menu to opposite ends */
  align-items: center;

  /* Positioning */
  position: fixed; /* Anchors the header to the top */
  top: 0;
  left: 0;
  width: 100%;
  z-index: 1000; /* Ensures it stays on top of other content */

  /* Styling */
  background-color: var(--primary-color);
  color: var(--text-on-primary);
  padding: 0 1.5rem; /* Adds horizontal space on the sides */
  height: 4rem; /* Define a fixed height */
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1); /* Optional: adds a subtle shadow */
}

.logo-container h1 {
  margin: 0;
  font-size: 1.5rem;
}

.main-nav {
  position: relative; /* Crucial for positioning the dropdown menu */
}

/* The <details> element acts as our dropdown container */
.menu-dropdown {
  cursor: pointer;
}

/* The <summary> is the visible "Menu" button */
.menu-dropdown summary {
  list-style: none; /* Removes the default triangle/marker */
  padding: 0.5rem 1rem;
  border: 1px solid var(--text-on-primary);
  border-radius: 4px;
  transition: background-color 0.2s ease-in-out;
}

.menu-dropdown summary:hover {
  background-color: rgba(0, 0, 0, 0.1);
}

/* Hide the default marker on Webkit browsers (Chrome, Safari) */
.menu-dropdown summary::-webkit-details-marker {
  display: none;
}

/* This is the list that appears when the menu is open */
.menu-items {
  position: absolute; /* Takes the list out of the normal document flow */
  top: calc(100% + 5px); /* Positions it just below the <summary> button */
  right: 0; /* Aligns to the right edge of the nav */

  /* Styling */
  background-color: var(--background-color);
  border: 1px solid #ddd;
  border-radius: 4px;
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
  list-style: none;
  padding: 0.5rem 0;
  margin: 0;
  min-width: 160px; /* Ensures the dropdown has a reasonable width */
  z-index: 1001;
}

/* Styling for each link inside the dropdown */
.menu-items li a {
  display: block;
  padding: 0.75rem 1.5rem;
  color: var(--text-on-secondary);
  text-decoration: none;
  cursor: pointer;
}

.menu-items li a:hover {
  background-color: oklch(0.9 0 0); /* A slightly darker gray on hover */
}
</style>
import { createRouter, createWebHistory } from 'vue-router'
import routes from '@/constants/routes'
import HomeView from '@/views/HomeView.vue'
import StatusView from '@/views/Status.vue'
import PlantView from '@/views/PlantView.vue'
import CultivarView from '@/views/CultivarView.vue'
import SpeciesView from '@/views/SpeciesView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: routes.ROUTE_INDEX,
      name: 'home',
      component: HomeView,
    },
    {path: routes.ROUTE_INTERNAL, name: 'internal', component: StatusView},
    {path: routes.ROUTE_PLANT_CULTIVAR, name: 'cultivar', component: CultivarView},
    {path: routes.ROUTE_PLANT_SPECIES, name: 'species', component: SpeciesView},
    {path: routes.ROUTE_PLANT, name: 'plant', component: PlantView},
  ],
})

export default router

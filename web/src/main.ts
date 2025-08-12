import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from '@/router';
import BaseAlert from '@/components/ui/BaseAlert.vue';
import BaseButton from '@/components/ui/BaseButton.vue';
import BaseCard from '@/components/ui/BaseCard.vue';
import BaseGraph from '@/components/ui/BaseGraph.vue';
import BaseMenu from '@/components/ui/BaseMenu.vue';
import BaseModal from '@/components/ui/BaseModal.vue';
import BaseTable from '@/components/ui/BaseTable.vue';
import BaseTabs from '@/components/ui/BaseTabs.vue';
import ExitButton from '@/components/ui/ExitButton.vue';
import App from './App.vue'

const app = createApp(App)

app.component('BaseAlert', BaseAlert)
app.component('BaseButton', BaseButton)
app.component('BaseCard', BaseCard)
app.component('BaseGraph', BaseGraph)
app.component('BaseMenu', BaseMenu)
app.component('BaseModal', BaseModal)
app.component('BaseTable', BaseTable)
app.component('BaseTabs', BaseTabs)
app.component('ExitButton', ExitButton)

app.use(createPinia())
app.use(router)

app.mount('#app')

const BASE_URL = import.meta.env.VITE_API_BASE_URL

if (!BASE_URL) {
    console.error("Required Environment variable VITE_API_BASE_URL is not set.")
}

const BTN_ADMIN = 'Admin'
const BTN_CLEAR = 'Clear';
const BTN_COPY = 'Copy';
const BTN_CREATE = 'Create';
const BTN_DELETE = 'Delete';
const BTN_DETAILS = 'Details';
const BTN_EDIT = 'Edit';
const BTN_HOME = 'Home';
const BTN_LOAD = 'Load';
const BTN_MENU = 'Menu';
const BTN_OFF = 'Off';
const BTN_ON = 'On';
const BTN_OPTIONS = 'Options';
const BTN_PIVOT = 'Pivot';
const BTN_SAVE = 'Save';

const BTN_PLANT = "Plants"
const BTN_PLANT_CULTIVAR = "Cultivars"
const BTN_PLANT_SPECIES = "Species"

const ID_PLANT_CULTIVAR_FORM = "plant-cultivar-form"
const ID_PLANT_FORM = "plant-form"
const ID_PLANT_SPECIES_FORM = "plant-species-form"

const TITLE_PLANT = "Plants"
const TITLE_PLANT_CULTIVAR = "Plant Cultivars"
const TITLE_PLANT_SPECIES = "Plant Species"

export default Object.freeze({
    BASE_URL,
    BTN_ADMIN,
    BTN_CLEAR,
    BTN_COPY,
    BTN_CREATE,
    BTN_DELETE,
    BTN_DETAILS,
    BTN_EDIT,
    BTN_HOME,
    BTN_LOAD,
    BTN_MENU,
    BTN_OFF,
    BTN_ON,
    BTN_OPTIONS,
    BTN_PIVOT,
    BTN_PLANT,
    BTN_PLANT_CULTIVAR,
    BTN_PLANT_SPECIES,
    BTN_SAVE,
    ID_PLANT_CULTIVAR_FORM,
    ID_PLANT_FORM,
    ID_PLANT_SPECIES_FORM,
    TITLE_PLANT,
    TITLE_PLANT_CULTIVAR,
    TITLE_PLANT_SPECIES,
});

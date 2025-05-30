

export type Button = {
  text: string
  onClick: Function
  classes?: Array<string>
};


export type Menu = {
  text: string
  buttons: Array<Button>
  classes?: Array<string>
};

export type Clickable = Button | Menu;


// meta note for my own thought process/understanding
//  FetchWrappable is a struct with containers for the following:
//    - the three attributes on UseFetchReturn<any>
//      - isFetching
//      - error
//      - data
//    - Slot props for each
//      - isFetchingProps
//      - error
//      - data

// can define sensible defaults at this level so pages don't break
// TODO: Implement the above below
export type FetchWrappable = {
  isFetching: any
  // isFetchingContent: any
  error: any
  // errorContent: any
  data: any
  // dataContent: any
}

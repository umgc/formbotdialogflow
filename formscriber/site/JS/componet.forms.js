Vue.component('forms', {
  props: {
    list: Array,
    filterKey: String
  },
  computed: { 
    filteredForms: function () {
      var filterKey = this.filterKey && this.filterKey.toLowerCase();
      var data = this.list
      if (filterKey) {
        data = data.filter(function (row) {
          return Object.keys(row).some(function (key) {
            return String(row[key]).toLowerCase().indexOf(filterKey) > -1;
          })
        })
      }
      return data;
    }
  },
  methods:{
    readForm(form){
      this.$emit('openformreader',form.id);
    },
    editForm(form){
      this.$emit('openformeditor',form.id);
    },
    copyTextArea(txt) {
        navigator.clipboard.writeText(txt);
    }
  },
  filters: {
    capitalize: function (str) {
      return str.charAt(0).toUpperCase() + str.slice(1)
    }
  },
  /* TODO
   * - Clean up layout
   * - Add toggel to show and hide the iframe
   */
  template: `  
  <article id="docs" class="list" style="float: left; width: 300px; height: 400px;">
    <input placeholder="Enter URL to Template Directory in Google Drive"></input>
    <ul>
      <li v-for="form in filteredForms">{{form.formName}}<div class="btn" @click="copyTextArea(form.URLl)">URL</div>
        <iframe src="https://docs.google.com/document/d/e/2PACX-1vSfv0ChJElCQbG0asDohdzZ90KetfqRf6jv7D3Vd8VHn3R5o5dHBgxqgkesGtQ3fnHvIdqrl8V-GcrJ/pub?embedded=true"></iframe>
      </li>
    </ul>
  </article>
      `
    })
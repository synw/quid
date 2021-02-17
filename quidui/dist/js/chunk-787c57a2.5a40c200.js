(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-787c57a2"],{"05a8":function(t,e,n){"use strict";n.r(e);var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("h1",{staticClass:"text-muted mt-3"},[t._v(" Orgs "),"addOrg"!==t.action?n("b-icon-plus",{staticClass:"mr-1",staticStyle:{color:"lightgrey"},on:{click:function(e){return t.$store.commit("action","addOrg")}}}):t._e()],1),t.state.isLoading?n("loading-indicator"):t._e(),n("div",[n("b-collapse",{staticClass:"mt-2",attrs:{id:"collapse-4"},model:{value:t.showActionBar,callback:function(e){t.showActionBar=e},expression:"showActionBar"}},["addOrg"===t.action?n("orgs-add",{on:{refresh:t.refresh}}):t._e()],1)],1),n("b-table",{staticClass:"mt-4",staticStyle:{"max-width":"650px"},attrs:{hover:"",bordeless:"",items:t.data,fields:t.fields},scopedSlots:t._u([{key:"cell(action)",fn:function(e){return[n("b-button",{attrs:{variant:"outline-danger"},on:{click:function(n){return t.confirmDeleteItem(e.item.id,e.item.name)}}},[t._v("Delete")])]}}])}),n("b-modal",{ref:"delete-modal",attrs:{title:"Delete org"},scopedSlots:t._u([{key:"modal-footer",fn:function(e){return[n("b-button",{attrs:{variant:"danger"},on:{click:function(e){return t.deleteItem(t.itemToDelete)}}},[t._v("Delete")]),n("b-button",{attrs:{variant:"warning"},on:{click:function(t){return e.cancel()}}},[t._v("Cancel")])]}}])},[t._v(" Delete "+t._s(t.itemToDelete.name)+"? ")])],1)},r=[],s=n("5530"),o=(n("96cf"),n("1da1")),i=n("2f62"),c=n("b758"),l=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("b-form",[n("b-card",{staticClass:"mb-2",staticStyle:{"max-width":"28rem"},attrs:{title:"Add an org"}},[n("b-card-text",[n("b-form-group",[n("b-form-input",{attrs:{state:t.isValidName,placeholder:"name"},model:{value:t.form.name,callback:function(e){t.$set(t.form,"name",e)},expression:"form.name"}}),n("b-form-invalid-feedback",{attrs:{state:t.isValidName}},[t._v("The name must be at least 4 characters long")])],1)],1),t.isValidName?n("b-button",{attrs:{variant:"success"},on:{click:t.postForm}},[t._v("Save")]):n("b-button",{attrs:{variant:"success",disabled:""}},[t._v("Save")]),n("b-button",{staticClass:"ml-2",attrs:{variant:"warning"},on:{click:function(e){return t.$store.commit("endAction")}}},[t._v("Cancel")])],1)],1)},u=[],m=(n("b0c0"),{data:function(){return{form:{name:""}}},methods:{postForm:function(){var t=this;return Object(o["a"])(regeneratorRuntime.mark((function e(){var n,a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,t.$api.post("/admin/orgs/add",{name:t.form.name});case 2:n=e.sent,a=n.error,null==a?(t.form.name="",t.$emit("refresh"),t.$store.commit("endAction")):409===a.reponse.status&&(t.form.name="");case 5:case"end":return e.stop()}}),e)})))()}},computed:{isValidName:function(){return 0===this.form.name.length?null:this.form.name.length>=4}}}),d=m,f=n("2877"),b=Object(f["a"])(d,l,u,!1,null,null,null),h=b.exports,g={components:{LoadingIndicator:c["a"],OrgsAdd:h},data:function(){return{data:[],state:{isLoading:!1},fields:[{key:"id",sortable:!0},{key:"name",sortable:!0},{key:"action",sortable:!1}],itemToDelete:{}}},methods:{fetchOrgs:function(){var t=this;return Object(o["a"])(regeneratorRuntime.mark((function e(){var n,a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return t.state.isLoading=!0,e.next=3,t.$api.get("/admin/orgs/all");case 3:n=e.sent,a=n.response,t.data=a.data,t.state.isLoading=!1;case 7:case"end":return e.stop()}}),e)})))()},confirmDeleteItem:function(t,e){this.itemToDelete={name:e,id:t},this.$refs["delete-modal"].show()},deleteItem:function(t){var e=this;return Object(o["a"])(regeneratorRuntime.mark((function n(){var a,r;return regeneratorRuntime.wrap((function(n){while(1)switch(n.prev=n.next){case 0:return e.$refs["delete-modal"].hide(),n.next=3,e.$api.post("/admin/orgs/delete",{id:t.id});case 3:a=n.sent,r=a.error,null===r&&(e.$bvToast.toast("Ok",{title:"Org deleted",autoHideDelay:1e3,variant:"success"}),e.fetchOrgs());case 6:case"end":return n.stop()}}),n)})))()},refresh:function(){this.fetchOrgs(),this.$bvToast.toast("ok",{title:"Org saved",variant:"success",autoHideDelay:1500})}},computed:Object(s["a"])(Object(s["a"])(Object(s["a"])({},Object(i["c"])(["action"])),Object(i["b"])({s:"showActionBar"})),{},{showActionBar:{get:function(){return this.s},set:function(t){return t}}}),mounted:function(){this.fetchOrgs()}},v=g,p=Object(f["a"])(v,a,r,!1,null,null,null);e["default"]=p.exports},b0c0:function(t,e,n){var a=n("83ab"),r=n("9bf2").f,s=Function.prototype,o=s.toString,i=/^\s*function ([^ (]*)/,c="name";a&&!(c in s)&&r(s,c,{configurable:!0,get:function(){try{return o.call(this).match(i)[1]}catch(t){return""}}})},b758:function(t,e,n){"use strict";var a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"text-center mt-5",staticStyle:{color:"lightgrey"}},[n("b-icon",{staticClass:"h1",attrs:{icon:"three-dots",animation:"cylon"}})],1)},r=[],s=n("2877"),o={},i=Object(s["a"])(o,a,r,!1,null,null,null);e["a"]=i.exports}}]);
//# sourceMappingURL=chunk-787c57a2.5a40c200.js.map
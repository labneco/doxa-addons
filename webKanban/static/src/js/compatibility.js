doxa.define('web_kanban.compatibility', function (require) {
"use strict";

var kanban_widgets = require('web_kanban.widgets');
var KanbanRecord = require('web_kanban.Record');
var KanbanColumn = require('web_kanban.Column');
var KanbanView = require('web_kanban.KanbanView');

return;
doxaerp = window.doxaerp || {};
doxaerp.web_kanban = doxaerp.web_kanban || {};
doxaerp.web_kanban.AbstractField = kanban_widgets.AbstractField;
doxaerp.web_kanban.KanbanGroup = KanbanColumn;
doxaerp.web_kanban.KanbanRecord = KanbanRecord;
doxaerp.web_kanban.KanbanView = KanbanView;

});

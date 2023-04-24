"use strict";
var IndexTable = (function () { 
    var menu = '<a href="#" class="btn btn-sm btn-light btn-active-light-primary" data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end">Actions<span class="svg-icon svg-icon-5 m-0"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M11.4343 12.7344L7.25 8.55005C6.83579 8.13583 6.16421 8.13584 5.75 8.55005C5.33579 8.96426 5.33579 9.63583 5.75 10.05L11.2929 15.5929C11.6834 15.9835 12.3166 15.9835 12.7071 15.5929L18.25 10.05C18.6642 9.63584 18.6642 8.96426 18.25 8.55005C17.8358 8.13584 17.1642 8.13584 16.75 8.55005L12.5657 12.7344C12.2533 13.0468 11.7467 13.0468 11.4343 12.7344Z" fill="currentColor" /></svg></span></a>'
    var t, table, acData, n = () => {
        t.querySelectorAll('[data-kt-ecommerce-product-filter="delete_row"]').forEach((t) => {
            if (t){
            t.addEventListener("click", function (t) {
                t.preventDefault();
                const n = t.target.closest("tr"),
                    r = n.querySelector('[data-kt-ecommerce-product-filter="name"]').innerText;
                    let id = n.querySelector('[data-kt-ecommerce-product-filter="delete_row"]').getAttribute('data-kt-ecommerce-id');
                    acData = "delete"
                    const res = actionManage(r, id, acData);
            });
            }
        });

        t.querySelectorAll('[data-kt-ecommerce-filter="ban_row"]').forEach((t) => {
            if (t){
            t.addEventListener("click",async function (t) {
                t.preventDefault();
                const n = t.target.closest("tr"),
                    r = n.querySelector('[data-kt-ecommerce-product-filter="name"]').innerText;
                    const id = n.querySelector('[data-kt-ecommerce-filter="ban_row"]').getAttribute('data-kt-ecommerce-id');
                    acData = "ban"
                    const res = await actionManage(r, id, acData);
            });
                
            }
        });

        t.querySelectorAll('[data-kt-ecommerce-filter="active_row"]').forEach((t) => {
            if (t){
            t.addEventListener("click",async function (t) {
                t.preventDefault();
                const n = t.target.closest("tr"),
                    r = n.querySelector('[data-kt-ecommerce-product-filter="name"]').innerText;
                    const id = n.querySelector('[data-kt-ecommerce-filter="active_row"]').getAttribute('data-kt-ecommerce-id');
                    acData = "active"
                    const res = await actionManage(r, id, acData);
            });
                
            }
        });
    };
    const actionManage = async (r, id, action) => {
        const result = await Swal.fire({
            text: "Are you sure you want to " + action + " " + r + "?",
            icon: "warning",
            showCancelButton: !0,
            buttonsStyling: !1,
            confirmButtonText: "Yes, "+action+"!",
            cancelButtonText: "No, cancel",
            customClass: { confirmButton: "btn fw-bold btn-danger", cancelButton: "btn fw-bold btn-active-light-primary" },
        });
        if (result.value) {
            let url = new URL(window.location);
            let pathname = url.pathname.replace(/\/$/, ''); // removes trailing slash if present
            try {
                const data = {
                    id: id,
                    action: action
                };
                const res = await axios.post('/api' + pathname + '/action', data, {
                    headers: {
                        "X-Requested-With": "XMLHttpRequest"
                    },
                    withCredentials: true
                });
                if(res.status == 204){
                    table.row($(n)).remove().draw();
                }
            } catch (error) {
                result.dismiss && Swal.fire({ text: r + " was not "+action+"ed.", icon: "error", buttonsStyling: !1, confirmButtonText: "Ok, got it!", customClass: { confirmButton: "btn fw-bold btn-primary" } });
            }
        } else {
            result.dismiss && Swal.fire({ text: r + " was not "+action+"ed.", icon: "error", buttonsStyling: !1, confirmButtonText: "Ok, got it!", customClass: { confirmButton: "btn fw-bold btn-primary" } });
        }
    }
    return {
        init: function () {
            (() => {
                t = document.querySelector("#kt_ecommerce_products_table")
                if (t) {
                let url = new URL(window.location)
                let pathname = url.pathname.replace(/\/$/, ''); // removes trailing slash if present
                // Get the number of columns
                var columnCount = $(t).find('thead th').length;
                // Calculate the index of the last column
                var secoundLastCol = columnCount - 2;
                
                table = $(t).DataTable({
                        serverSide: true,
                        processing: true,
                        info: 1,
                        ajax: {
                            url: "/api" + pathname,
                            type: 'POST',
                            beforeSend: function (request) {
                                request.setRequestHeader("Accept", "dataTables/json");
                            },
                            data: function (d) {
                                let nm = document.querySelector('[data-kt-ecommerce-product-filter="status"]');
                                if(nm) {
                                    d.status = nm.value;
                                }
                                if(d.draw == 1){
                                    d.order[0]['column'] = secoundLastCol;
                                    d.order[0]['dir'] = 'desc';
                                }
                            }
                        },
                        order: [[secoundLastCol, 'desc']],
                        lengthMenu: [ [10, 25, 50, 100, 500, -1], ["Show 10", "Show 25", "Show 50", "Show 100", "Show 500", "Show All"] ],
                        dom: 'rtilBp',
                        buttons: [
                            {
                                extend: 'excelHtml5',
                                text: 'Export to Excel',
                                className: 'btn btn-sm btn-light-primary',
                                exportOptions: {
                                    columns: function (idx, data, node) {
                                        return node.textContent.toLowerCase() !== 'action';
                                    }
                                },
                            },
                            {
                                extend: 'csvHtml5',
                                text: 'CSV',
                                className: 'btn btn-sm btn-light-primary',
                                exportOptions: {
                                    columns: function (idx, data, node) {
                                        return node.textContent.toLowerCase() !== 'action';
                                    }
                                },
                            },
                            {
                                extend: 'pdfHtml5',
                                text: 'PDF',
                                className: 'btn btn-sm btn-light-primary',
                                exportOptions: {
                                    columns: function (idx, data, node) {
                                        return node.textContent.toLowerCase() !== 'action';
                                    }
                                },
                                orientation: 'landscape',
                                customize: function (doc) {
                                    doc.defaultStyle.fontSize = 12;
                                    doc.content[1].table.widths = Array(doc.content[1].table.body[0].length + 1).join('*').split('');
                                    
                                    // Add "Showing X to X of X records" to each page
                                    doc['footer'] = (function (page, pages) {
                                        return {
                                            columns: [
                                                {
                                                    alignment: 'center',
                                                    text: ['Page ', { text: page.toString() }, ' of ', { text: pages.toString() }, ' | ', { text: 'Total records ' + (doc.content[1].table.body.length - 1) }]
                                                }
                                            ],
                                            margin: [10, 0]
                                        }
                                    });
                                }
                            },
                        ],
                        paging: true,
                        search: false,
                        columnDefs: [
                            {
                                targets: -1,
                                orderable: false,
                                render: function (data, type, row) {

                                    if (data.includes('{"action":')){
                                        
                                        let d = JSON.parse(data)
                                        let html = '';
                                        let target = '';
                                        d.action.forEach(item => {
                                            let filter = '';
                                            if (item.filter_delete) {
                                                filter = ' data-kt-ecommerce-product-filter="'+item.filter_delete+'"' + ' data-kt-ecommerce-id="'+item.id+'"';
                                            }
                                            if (item.filter_ban) {
                                                filter = ' data-kt-ecommerce-filter="'+item.filter_ban+'"' + ' data-kt-ecommerce-id="'+item.id+'"';
                                            }
                                            if (item.filter_active) {
                                                filter = ' data-kt-ecommerce-filter="'+item.filter_active+'"' + ' data-kt-ecommerce-id="'+item.id+'"';
                                            }
                                            if (item.target) {
                                                target = ' target="'+item.target+'"';
                                            }
                                            html += '<div class="menu-item px-3"><a href="'+ item.url+'" class="menu-link px-3"' + filter + target + '>'+item.name+'</a></div>';
                                        })
                                        return menu + '<div class="menu menu-sub menu-sub-dropdown menu-column menu-rounded menu-gray-600 menu-state-bg-light-primary fw-semibold fs-7 w-125px py-4" data-kt-menu="true">' + html + '</div>';
                                    } else{
                                        return '<a href="' + data + '">' + data + '</a>';
                                    }
                                }
                            }
                        ],
                        createdRow: function (row, data, dataIndex) {
                            $('td', row).eq(0).attr('data-kt-ecommerce-product-filter', 'name');
                        },
                        initComplete: function (settings, json) {
                            // Re-initialize menu
                            KTMenu.createInstances();
                        },
                });
                
                table.on("draw", (t) => {
                    n();
                    KTMenu.createInstances();
                });
                
                // table.on("page", (t) => {
                //     let url = new URL(window.location);
                //     var pageInfo = table.page.info();
                //     history.pushState({}, '', url.pathname + '?page=' + (pageInfo.page+1));
                // });

                document.querySelector('[data-kt-ecommerce-product-filter="search"]').addEventListener("keyup", (t) => {
                    table.search(t.target.value).draw();
                });

                }
                (() => {
                    const tb = document.querySelector('[data-kt-ecommerce-product-filter="status"]');
                    $(tb).on("change", (t) => {
                        let n = t.target.value;
                        "all" === n && (n = ""), table.column(6).search(n).draw();
                        console.log("change",n);
                    });
                })(),
                n();
            })();
        }
    }
})();
KTUtil.onDOMContentLoaded(function () {
    IndexTable.init();
});
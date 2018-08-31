package shell

import (
	"fmt"

	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
	"yunion.io/x/onecloud/pkg/mcclient/options"
)

func init() {
	R(&options.ServerListOptions{}, "server-list", "List virtual servers", func(s *mcclient.ClientSession, opts *options.ServerListOptions) error {
		params, err := options.ListStructToParams(opts)
		if err != nil {
			return err
		}
		result, err := modules.Servers.List(s, params)
		if err != nil {
			return err
		}
		printList(result, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerShowOptions{}, "server-show", "Show details of a server", func(s *mcclient.ClientSession, opts *options.ServerShowOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		result, err := modules.Servers.Get(s, opts.ID, params)
		if err != nil {
			return err
		}
		printObject(result)
		return nil
	})

	R(&options.ServerIdOptions{}, "server-metadata", "Show metadata of a server", func(s *mcclient.ClientSession, opts *options.ServerIdOptions) error {
		result, err := modules.Servers.GetMetadata(s, opts.ID, nil)
		if err != nil {
			return err
		}
		printObject(result)
		return nil
	})

	R(&options.ServerCreateOptions{}, "server-create", "Create a server", func(s *mcclient.ClientSession, opts *options.ServerCreateOptions) error {
		params, err := opts.Params()
		if err != nil {
			return err
		}
		count := options.IntV(opts.Count)
		if options.BoolV(opts.DryRun) {
			results, err := modules.SchedManager.DoScheduleListResult(s, params, count)
			if err != nil {
				return err
			}
			printList(results, []string{"id", "name", "rank", "capacity", "error"})
		} else {
			taskNotify := options.BoolV(opts.TaskNotify)
			if taskNotify {
				s.PrepareTask()
			}
			if count > 1 {
				results := modules.Servers.BatchCreate(s, params, count)
				printBatchResults(results, modules.Servers.GetColumns(s))
			} else {
				server, err := modules.Servers.Create(s, params)
				if err != nil {
					return err
				}
				printObject(server)
			}
			if taskNotify {
				s.WaitTaskNotify()
			}
		}
		return nil
	})

	R(&options.ServerIdOptions{}, "server-logininfo", "Get login info of a server", func(s *mcclient.ClientSession, opts *options.ServerIdOptions) error {
		srvid, e := modules.Servers.GetId(s, opts.ID, nil)
		if e != nil {
			return e
		}
		i, e := modules.Servers.GetLoginInfo(s, srvid, nil)
		if e != nil {
			return e
		}
		printObject(i)
		return nil
	})

	R(&options.ServerIdsOptions{}, "server-start", "Start servers", func(s *mcclient.ClientSession, opts *options.ServerIdsOptions) error {
		ret := modules.Servers.BatchPerformAction(s, opts.ID, "start", nil)
		printBatchResults(ret, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerIdsOptions{}, "server-syncstatus", "Sync servers status", func(s *mcclient.ClientSession, opts *options.ServerIdsOptions) error {
		ret := modules.Servers.BatchPerformAction(s, opts.ID, "syncstatus", nil)
		printBatchResults(ret, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerIdsOptions{}, "server-sync", "Sync servers configurations", func(s *mcclient.ClientSession, opts *options.ServerIdsOptions) error {
		ret := modules.Servers.BatchPerformAction(s, opts.ID, "sync", nil)
		printBatchResults(ret, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerStopOptions{}, "server-stop", "Stop servers", func(s *mcclient.ClientSession, opts *options.ServerStopOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		ret := modules.Servers.BatchPerformAction(s, opts.ID, "stop", params)
		printBatchResults(ret, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerIdsOptions{}, "server-suspend", "Suspend servers", func(s *mcclient.ClientSession, opts *options.ServerIdsOptions) error {
		ret := modules.Servers.BatchPerformAction(s, opts.ID, "suspend", nil)
		printBatchResults(ret, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerResetOptions{}, "server-reset", "Reset servers", func(s *mcclient.ClientSession, opts *options.ServerResetOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		ret := modules.Servers.BatchPerformAction(s, opts.ID, "reset", params)
		printBatchResults(ret, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerIdsOptions{}, "server-purge", "Purge obsolete servers", func(s *mcclient.ClientSession, opts *options.ServerIdsOptions) error {
		ret := modules.Servers.BatchPerformAction(s, opts.ID, "purge", nil)
		printBatchResults(ret, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerDeleteOptions{}, "server-delete", "Delete servers", func(s *mcclient.ClientSession, opts *options.ServerDeleteOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		ret := modules.Servers.BatchDeleteWithParam(s, opts.ID, params, nil)
		printBatchResults(ret, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerIdsOptions{}, "server-cancel-delete", "Cancel pending delete servers", func(s *mcclient.ClientSession, opts *options.ServerIdsOptions) error {
		ret := modules.Servers.BatchPerformAction(s, opts.ID, "cancel-delete", nil)
		printBatchResults(ret, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerIdOptions{}, "server-vnc", "Show vnc info of server", func(s *mcclient.ClientSession, opts *options.ServerIdOptions) error {
		ret, e := modules.Servers.GetSpecific(s, opts.ID, "vnc", nil)
		if e != nil {
			return e
		}
		printObject(ret)
		return nil
	})

	R(&options.ServerIdOptions{}, "server-desc", "Show desc info of server", func(s *mcclient.ClientSession, opts *options.ServerIdOptions) error {
		ret, e := modules.Servers.GetSpecific(s, opts.ID, "desc", nil)
		if e != nil {
			return e
		}
		printObject(ret)
		return nil
	})

	R(&options.ServerIdOptions{}, "server-status", "Show status of server", func(s *mcclient.ClientSession, opts *options.ServerIdOptions) error {
		ret, e := modules.Servers.GetSpecific(s, opts.ID, "status", nil)
		if e != nil {
			return e
		}
		printObject(ret)
		return nil
	})

	R(&options.ServerUpdateOptions{}, "server-update", "Update servers", func(s *mcclient.ClientSession, opts *options.ServerUpdateOptions) error {
		params, err := opts.Params()
		if err != nil {
			return err
		}
		if params.Size() == 0 {
			return InvalidUpdateError()
		}
		result := modules.Servers.BatchPut(s, opts.ID, params)
		printBatchResults(result, modules.Servers.GetColumns(s))
		return nil
	})

	R(&options.ServerSendKeyOptions{}, "server-send-keys", "Send keys to server", func(s *mcclient.ClientSession, opts *options.ServerSendKeyOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		srv, err := modules.Servers.PerformAction(s, opts.ID, "sendkeys", params)
		if err != nil {
			return err
		}
		printObject(srv)
		return nil
	})

	R(&options.ServerDeployOptions{}, "server-deploy", "Deploy hostname and keypair to a stopped virtual server", func(s *mcclient.ClientSession, opts *options.ServerDeployOptions) error {
		params, err := opts.Params()
		if err != nil {
			return err
		}
		srv, e := modules.Servers.PerformAction(s, opts.ID, "deploy", params)
		if e != nil {
			return e
		}
		printObject(srv)
		return nil
	})

	R(&options.ServerSecGroupOptions{}, "server-assign-secgroup", "Assign security group to a VM", func(s *mcclient.ClientSession, opts *options.ServerSecGroupOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		srv, e := modules.Servers.PerformAction(s, opts.ID, "assign-secgroup", params)
		if e != nil {
			return e
		}
		printObject(srv)
		return nil
	})

	R(&options.ServerSecGroupOptions{}, "server-assign-admin-secgroup", "Assign administrative security group to a VM", func(s *mcclient.ClientSession, opts *options.ServerSecGroupOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		srv, e := modules.Servers.PerformAction(s, opts.ID, "assign-admin-secgroup", params)
		if e != nil {
			return e
		}
		printObject(srv)
		return nil
	})

	R(&options.ServerIdOptions{}, "server-revoke-secgroup", "Assign security group to a VM", func(s *mcclient.ClientSession, opts *options.ServerIdOptions) error {
		srv, e := modules.Servers.PerformAction(s, opts.ID, "revoke-secgroup", nil)
		if e != nil {
			return e
		}
		printObject(srv)
		return nil
	})

	R(&options.ServerIdOptions{}, "server-revoke-admin-secgroup", "Assign administrative security group to a VM", func(s *mcclient.ClientSession, opts *options.ServerIdOptions) error {
		srv, e := modules.Servers.PerformAction(s, opts.ID, "revoke-admin-secgroup", nil)
		if e != nil {
			return e
		}
		printObject(srv)
		return nil
	})

	R(&options.ServerMonitorOptions{}, "server-monitor", "Send commands to qemu monitor", func(s *mcclient.ClientSession, opts *options.ServerMonitorOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		ret, err := modules.Servers.GetSpecific(s, opts.ID, "monitor", params)
		if err != nil {
			return err
		}
		result, err := ret.GetString("results")
		if err != nil {
			return err
		}
		fmt.Println(result)
		return nil
	})

	R(&options.ServerSaveImageOptions{}, "server-save-image", "Save root disk to new image and upload to glance.", func(s *mcclient.ClientSession, opts *options.ServerSaveImageOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		srv, err := modules.Servers.PerformAction(s, opts.ID, "save-image", params)
		if err != nil {
			return err
		}
		printObject(srv)
		return nil
	})

	type ServerChangeOwnerOptions struct {
		ID      string `help:"Server to change owner"`
		PROJECT string `help:"Project ID or change"`
		RawId   bool   `help:"User raw ID, instead of name"`
	}
	R(&ServerChangeOwnerOptions{}, "server-change-owner", "Change owner porject of a server", func(s *mcclient.ClientSession, opts *ServerChangeOwnerOptions) error {
		params := jsonutils.NewDict()
		if opts.RawId {
			projid, err := modules.Projects.GetId(s, opts.PROJECT, nil)
			if err != nil {
				return err
			}
			params.Add(jsonutils.NewString(projid), "tenant")
			params.Add(jsonutils.JSONTrue, "raw_id")
		} else {
			params.Add(jsonutils.NewString(opts.PROJECT), "tenant")
		}
		srv, err := modules.Servers.PerformAction(s, opts.ID, "change-owner", params)
		if err != nil {
			return err
		}
		printObject(srv)
		return nil
	})

	R(&options.ServerRebuildRootOptions{}, "server-rebuild-root", "Rebuild VM root image with new template", func(s *mcclient.ClientSession, opts *options.ServerRebuildRootOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		srv, err := modules.Servers.PerformAction(s, opts.ID, "rebuild-root", params)
		if err != nil {
			return err
		}
		printObject(srv)
		return nil
	})

	R(&options.ServerChangeConfigOptions{}, "server-change-config", "Change configuration of VM", func(s *mcclient.ClientSession, opts *options.ServerChangeConfigOptions) error {
		params, err := options.StructToParams(opts)
		if err != nil {
			return err
		}
		if params.Size() == 0 {
			return InvalidUpdateError()
		}
		srv, err := modules.Servers.PerformAction(s, opts.ID, "change-config", params)
		if err != nil {
			return err
		}
		printObject(srv)
		return nil
	})

	type ServerInsertISOOptions struct {
		ID  string `help:"server ID or Name"`
		ISO string `help:"Glance image ID of the ISO"`
	}
	R(&ServerInsertISOOptions{}, "server-insert-iso", "Insert an ISO image into server's cdrom", func(s *mcclient.ClientSession, opts *ServerInsertISOOptions) error {
		img, err := modules.Images.Get(s, opts.ISO, nil)
		if err != nil {
			return err
		}
		imgId, err := img.GetString("id")
		if err != nil {
			return err
		}
		params := jsonutils.NewDict()
		params.Add(jsonutils.NewString(imgId), "image_id")
		result, err := modules.Servers.PerformAction(s, opts.ID, "insertiso", params)
		if err != nil {
			return err
		}
		printObject(result)
		return nil
	})

	R(&options.ServerIdOptions{}, "server-eject-iso", "Eject iso from servers' cdrom", func(s *mcclient.ClientSession, opts *options.ServerIdOptions) error {
		result, err := modules.Servers.PerformAction(s, opts.ID, "ejectiso", nil)
		if err != nil {
			return err
		}
		printObject(result)
		return nil
	})

	R(&options.ServerIdOptions{}, "server-iso", "Show server's mounting ISO information", func(s *mcclient.ClientSession, opts *options.ServerIdOptions) error {
		results, err := modules.Servers.GetSpecific(s, opts.ID, "iso", nil)
		if err != nil {
			return err
		}
		printObject(results)
		return nil
	})
}

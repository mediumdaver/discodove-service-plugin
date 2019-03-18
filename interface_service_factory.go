/* Written by Dave Richards.
 *
 * This is the top-level plugin interface, where you produce a new instance of the service.
 */
package discodove_interface_service

import (
	"log/syslog"
)

type DiscoDoveServiceFactory interface {

	/* Initialise this service, if necessary.  This will be called before the first use of this service,
	 * if you feel compelled to set something up, perhaps a control/query/admin thread or something, then
	 * do it here.  
	 *
	 * Each plugin is responsible for it's own logging, i suggest syslog, but it's your call.  I was going
	 * to pass in a *syslog.Writer but it has a mutex in there, and i don't want the service threads to be
	 * blocking on writing to syslog - so you need to scale logging yourself.
	 * 
	 * We use Viper for config - so you can access the discodove config too by using viper, so feel free to
	 * include your own config directive in there, under it's own section.
	 *
	 * name	 	: will be the name of the process, in 99.999% of cases it will just be "discodove" - please
	 *            prefix your log messages with this.
	 * syslogFacility : which facility to use in syslog, if that's how you are logging - otherwise ignore it.
	 */
	NewService(name string, syslogFacility syslog.Priority) DiscoDoveService
}
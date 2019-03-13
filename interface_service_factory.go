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
	 * name	 	: will be the name of the process, in 99.999% of cases it will just be "discodove"
	 * logger	: a handle to a syslog.Writer to write your log messages to
	 */
	NewService(string, *syslog.Writer) DiscoDoveService
}
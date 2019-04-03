/* Written by Dave Richards.
 *
 * This is the top-level plugin interface, where you produce a new instance of the service.
 */
package discodove_interface_service

import (
	"log/syslog"
	"github.com/spf13/viper"
	"net"
	"github.com/mediumdaver/discodove-auth-plugin"
)
 
type DiscoDoveServiceFactory interface {

	/* This will be called once when we load this module a service instance, if you feel compelled to set something 
	 * up, perhaps a control/query/admin thread or something, then do it here in a controlled manner - similarly if 
	 * you want to pool connections, etc....  
	 *
	 * Each plugin is responsible for creating it's own syslog connection as *syslog.Writer has a mutex, and 
	 * I don't want the service threads to be blocking on writing to syslog - so you need to scale logging yourself.
	 * 
	 * We use Viper for config, and i will pass in the config directives for your module, but as it's viper you
	 * can access the entire discodove config too.  Feel free to specify your own config directives.
	 *
	 * name	 	: will be the name of the process, in 99.999% of cases it will just be "discodove" - please
	 *            prefix your log messages with this and perhaps your own identifier e.g. "imap" or "pop3".
	 * syslogFacility : which facility to use in syslog.
	 * conf: a Viper subtree configuration for this service as specified in the discodove config.
	 */
	Initialize(name string, syslogFacility syslog.Priority, auth discodove_interface_auth.DiscoDoveAuthService) error

	/* Because we can have a number of instances of the same service, we need seperate instances such that they may
	 * maintain seperate configurations. e.g. imap v's imaps.  We call this function to produce a new instance
	 * of the service.
	 */
	NewService(conf *viper.Viper) DiscoDoveServicePlugin
}

type DiscoDoveServicePlugin interface {

	/* The real work is done here, this method will be called, in context of a go routine, for every 
	 * connection made on the listen port.  DiscoDove will not attempt to validate the connection in any
	 * way, the assumption is if it passed all the firewalls and discodove's controls (SSL for example),
	 * then it's over to you.
	 * 
	 * Please close the connection and exit when you're done.
	 *
	 * conn 			: the connection.
	 * authenticators	: a collection of authenticators to call in priority order.
	 * stores			: all stores that have been initiated for use.
	 * secure			: a method of encrypting a plain text session if needed, e.g. STARTTLS
	 */
    HandleConnection(c net.Conn, alreadyEncrypted bool)
}